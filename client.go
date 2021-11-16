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

	d "github.com/kaeppen/disys-mandatory2/dimutex"
	"google.golang.org/grpc"
)

type diMutexClient struct {
	d.UnimplementedDiMutexServer
	state     State
	timestamp int32
	ctx       context.Context
	peers     map[int]d.DiMutexClient
	name      string
	id        int32
	queue     []int32
	replies   int32
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
	id, _ := strconv.Atoi(os.Getenv("ID"))
	c.id = int32(id)
	c.state = Released
	c.timestamp = 0
	ctx, cancel := context.WithCancel(context.Background())
	c.ctx = ctx
	defer cancel()

	//register the node as a server
	server := grpc.NewServer()
	d.RegisterDiMutexServer(server, &c)

	port := c.id + 8080
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Printf("Failed to listen on port %v", listen.Addr())
	}
	go serve(server, listen, &c)

	//set up connection to peers
	setupConnection(&c)
	c.participate()
}

func (c *diMutexClient) participate() {
	for {
		//let one node have the advantage to avoid deadlock
		if c.name != "Mikki" {
			time.Sleep(1 * time.Second)
		}
		request := d.Request{Message: fmt.Sprintf("%v (node %v) want's the critical section!", c.name, c.id), Lamport: c.timestamp, Id: c.id}
		c.RequestAccess(c.ctx, &request)

		//as long as we don't have enough replies
		for int(c.replies) < len(c.peers) {
			//waiting..
		}
		//we've got N-1 replies and can now access the CS
		c.HoldAndRelease(c.ctx, &d.Empty{})

		//now i'll reply to my queue
		c.ReplyToQueue()

	}
}

func demo() {
	//create some demo behavior here (for the logs)
}

func serve(server *grpc.Server, listener net.Listener, c *diMutexClient) {
	log.Printf("Server listening on %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Printf("Node %v failed to serve: %v", c.name, err)
	}
}

func (c *diMutexClient) Hello(ctx context.Context, in *d.Empty) (*d.Empty, error) {
	log.Print("HEJSA :)))")

	return &d.Empty{}, nil
}

//local actions -> OVERVEJ NAVNE
func GetAccess(message string, c *diMutexClient) {
	c.timestamp++    //bump logical clock
	c.state = Wanted //set the state to wanted
	request := d.Request{Message: "I want access!", Lamport: int32(c.timestamp), Id: c.id}
	c.RequestAccess(c.ctx, &request)
}

//The "multicast" part of the algorithm
func (c *diMutexClient) RequestAccess(ctx context.Context, in *d.Request) (*d.Empty, error) {
	log.Printf("%v (node %v) requesting access to cs", c.name, c.id)
	//bump own clock before sending out the message
	c.timestamp++
	//set own state to wanted
	c.state = Wanted

	//tell my peers that i want the critical section
	peers := c.peers
	for i := range c.peers {
		_, err := peers[i].AnswerRequest(ctx, in)
		if err != nil {
			log.Fatalf("Error while getting replies %v", err)
		}
	}
	return &d.Empty{}, nil
}

func (c *diMutexClient) HoldAndRelease(ctx context.Context, empty *d.Empty) (*d.Reply, error) {
	log.Printf("%v (node %v) has gotten access to cs", c.name, c.id)
	c.state = Held
	//Hold the critical section for 7 seconds
	time.Sleep(7 * time.Second)
	//Release it
	c.state = Released
	log.Printf("%v (node %v) has released cs", c.name, c.id)
	return &d.Reply{}, nil
}

func (c *diMutexClient) ReplyToQueue() {
	for i := range c.queue {
		//reply := d.Reply{Message: fmt.Sprintf("Node %v replying to node %v's request", c.id, c.queue[i]), Lamport: c.timestamp, Id: c.id}
		log.Printf("%v (node %v) replying to node %v's request", c.name, c.id, c.queue[i])
		c.peers[int(c.queue[i])].Grant(c.ctx, &d.Empty{})
		//reset queue size
	}
	c.queue = make([]int32, 0)
}

func (c *diMutexClient) GrantAccess(ctx context.Context, id int32) {
	_, err := c.peers[int(id)].Grant(ctx, &d.Empty{})
	if err != nil {
		log.Fatalf("Error while granting access to node %v: %v", id, err)
	}
}

func (c *diMutexClient) Grant(ctx context.Context, reply *d.Empty) (*d.Empty, error) {
	c.replies++
	return &d.Empty{}, nil
}

func hasPrecedence(ownTime int32, ownId int32, otherTime int32, otherId int32) bool {
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

func (c *diMutexClient) AnswerRequest(ctx context.Context, request *d.Request) (*d.Empty, error) {
	c.timestamp++ //increment before doing anything
	//if this node already has access or wants it and also has "more right"
	if c.state == Held || (c.state == Wanted && hasPrecedence(c.timestamp, c.id, request.Lamport, request.Id)) {
		//queue the request from other node
		c.queue = append(c.queue, request.Id)
	} else {
		//else, just grant access
		//reply := d.Reply{Message: fmt.Sprintf("Node %v replying to node %v's request", c.id, request.Id), Lamport: c.timestamp, Id: c.id}
		log.Printf("%v (node %v) replying to node %v's request", c.name, c.id, request.Id)
		c.GrantAccess(ctx, request.Id)
	}
	return &d.Empty{}, nil
}

//Set up connection to peers
func setupConnection(c *diMutexClient) {
	c.peers = make(map[int]d.DiMutexClient)
	for i := 0; i < 4; i++ {
		port := fmt.Sprintf("dimutex_%v:%v", i+1, i+8080)
		if int32(i) != c.id {
			conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("Could not connect: %s", err)
			}
			c.peers[i] = d.NewDiMutexClient(conn)
			log.Printf("node %v added node %v to list of peers", c.id, i)
		}
	}
}

//doesn't get called yet
func updateLamport(own int32, other int32) int32 {
	if own >= other {
		return own + 1
	}
	return other + 1
}
