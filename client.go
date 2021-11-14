package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/serf/serf"
	d "github.com/kaeppen/disys-mandatory2/dimutex"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type diMutexClient struct {
	cluster   *serf.Serf
	state     State
	timestamp int
	ctx       context.Context
	peers     []d.DiMutexClient //bliver det et problem med denne type som jo ikke har cluster, state osv.?
}

const (
	port = "???"
)

func main() {
	//These addresses work with the dockerfile from the example
	cluster, err := setupCluster(
		os.Getenv("ADVERTISE_ADDR"),
		os.Getenv("CLUSTER_ADDR"))
	if err != nil {
		log.Fatal(err)
	}
	defer cluster.Leave()

	c := diMutexClient{}
	c.state = Released
	c.timestamp = 0
	c.ctx = context.Background()
	waiter := time.Tick(2 * time.Second)
	//ctx := context.Background()
	select {
	case <-waiter:
		setupConnection(&c)
		for {
			fmt.Printf("Hello world!")
		}
		//ctx, _ := context.WithTimeout(ctx, time.Second*2)

	}
	//setupConnection(&c)
	//for {
	//	fmt.Printf("Hello world!")
	//}

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
	c.state = Held
	//Hold the critical section for 7 seconds
	time.Sleep(7 * time.Second)
	//Release it
	c.state = Released

	//maybe broadcast

	return nil //wrong?
}

func hasPrecedence(own int, recieved int) bool {
	return own > recieved
}

func (c *diMutexClient) AnswerRequest(ctx context.Context, request *d.AccessRequest) (*d.RequestAnswer, error) {
	c.timestamp++ //increment before doing anything
	//if this node already has access or wants it and also has "more right"
	if c.state == Held || (c.state == Wanted && hasPrecedence(c.timestamp, int(request.Lamport))) {
		//queue the request from other node
		return nil, nil //return value?
	} else {
		//else, just send a reply
		answer := &d.RequestAnswer{}
		return answer, nil
	}
}

//"converts" serf members to grpc clients
func setupConnection(c *diMutexClient) {
	members := getOtherMembers(c.cluster)
	for i := 0; i < len(members); i++ {
		addr := members[i].Addr.String()
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)

		}
		c.peers[i] = d.NewDiMutexClient(conn)
	}

}

//nakket fra eksempel
func getOtherMembers(cluster *serf.Serf) []serf.Member {
	members := cluster.Members()
	for i := 0; i < len(members); {
		if members[i].Name == cluster.LocalMember().Name || members[i].Status != serf.StatusAlive {
			if i < len(members)-1 {
				members = append(members[:i], members[i+1:]...)
			} else {
				members = members[:i]
			}
		} else {
			i++
		}
	}
	return members
}
