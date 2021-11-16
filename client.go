package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/hashicorp/serf/serf"
	d "github.com/kaeppen/disys-mandatory2/dimutex"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Server struct {
	d.UnimplementedDiMutexServer
}

type diMutexClient struct {
	state     State
	timestamp int
	ctx       context.Context
	peers     []d.DiMutexClient //bliver det et problem med denne type som jo ikke har cluster, state osv.?
	name      string
	id        int
	server    *grpc.Server
}

func main() {
	//set up logging
	//os.Remove("../Logfile.txt") //Delete the file to ensure a fresh log for every session
	f, erro := os.OpenFile("./Logfile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if erro != nil {
		log.Fatalf("Logfile error")
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	c := diMutexClient{}
	c.name = os.Getenv("NAME")
	c.id, _ = strconv.Atoi(os.Getenv("ID"))
	c.state = Released
	c.timestamp = 0
	c.ctx = context.Background()
	waiter := time.Tick(2 * time.Second)

	//register the node as a server
	server := grpc.NewServer()
	d.RegisterDiMutexServer(server, &Server{})
	c.server = server

	//listen (attempt to use serf agent port -> might be wrong)¨
	//fmt.Printf(":%v", strconv.Itoa((int(c.cluster.LocalMember().Port))))
	port := c.id + 8080
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Printf("Failed to listen on port %v", listen.Addr())
	}
	log.Printf("Sserver listening on %v", listen.Addr())
	if err := server.Serve(listen); err != nil {
		log.Printf("Node %v failed to serve: %v", c.name, err)
	}

	select {
	case <-waiter:
		setupConnection(&c)
	}
	c.peers[0].Hello(c.ctx, &d.Empty{})
	//for {
	//	select {
	//	case <-waiter:
	//		request := &d.AccessRequest{Message: "Hey", Lamport: int32(c.timestamp), Id: 9000} //bogus id
	//		c.RequestAccess(c.ctx, request)
	//	}
	//}
}

func (c *diMutexClient) Hello(in *d.Empty) {
	log.Print("HEJSA :)))")
}

func setupCluster(advertiseAddr string, clusterAddr string) (*serf.Serf, error) {
	conf := serf.DefaultConfig()
	conf.Init()
	conf.MemberlistConfig.AdvertiseAddr = advertiseAddr
	cluster, err := serf.Create(conf)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create cluster")
	}

	_, err = cluster.Join([]string{clusterAddr}, true)
	if err != nil {
		log.Printf("Couldn't join cluster, starting own: %v\n", err)
	}

	return cluster, nil
}

//local actions -> OVERVEJ NAVNE
func GetAccess(message string, c *diMutexClient) {
	c.timestamp++    //bump logical clock
	c.state = Wanted //set the state to wanted
	request := &d.AccessRequest{Message: "I want access!", Lamport: int32(c.timestamp), Id: 9000}
	c.RequestAccess(c.ctx, request)
}

//The "multicast" part of the algorithm
func (c *diMutexClient) RequestAccess(ctx context.Context, in *d.AccessRequest, opts ...grpc.CallOption) (*d.AccessGrant, error) {
	log.Printf("%v requesting access to cs", c.name)
	//bump own clock before sending out the message
	c.timestamp++
	//set own state to wanted
	c.state = Wanted

	replies := 0
	peers := c.peers
	for i := 0; i < len(peers); i++ {
		answer, err := peers[i].AnswerRequest(ctx, in)
		if answer != nil && err != nil {
			replies++
		}
	}
	//if i receive N-1 replies, then i can have the critical section
	if replies == len(peers) {
		c.HoldAndRelease(ctx, &d.Empty{})
	}

	//compiler gets happy -> revisit return
	return nil, nil
}

func (c *diMutexClient) HoldAndRelease(ctx context.Context, empty *d.Empty) *d.Empty {
	log.Printf("%v has gotten access to cs", c.name)
	c.state = Held
	//Hold the critical section for 7 seconds
	time.Sleep(7 * time.Second)
	//Release it
	c.state = Released
	log.Printf("%v has released cs", c.name)

	//maybe broadcast

	return empty
}

func hasPrecedence(ownTime int, ownId int, otherTime int, otherId int) bool {
	if ownTime > otherTime {
		return true
	} else if otherTime > ownTime {
		return false
	} else {
		//if timestamps are equal, tiebreaker is node id
		if ownId > otherId {
			return true
		} else {
			return false
		}
	}
}

func (c *diMutexClient) AnswerRequest(ctx context.Context, request *d.AccessRequest) (*d.RequestAnswer, error) {
	log.Print("Jeg er lige blevet ringet op med et gRPC kald, av av")
	c.timestamp++ //increment before doing anything
	//if this node already has access or wants it and also has "more right"
	if c.state == Held || (c.state == Wanted && hasPrecedence(c.timestamp, c.id, int(request.Lamport), int(request.Id))) {
		//queue the request from other node
		return nil, nil //return value?
	} else {
		//else, just send a reply
		answer := &d.RequestAnswer{}
		return answer, nil
	}
}

//Set up connection to peers
func setupConnection(c *diMutexClient) {
	//harcode connection to peers
	for i := 0; i < 4; i++ {
		port := fmt.Sprintf(":&v", c.id+8080)
		conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		c.peers = append(c.peers, d.NewDiMutexClient(conn))
	}
}
