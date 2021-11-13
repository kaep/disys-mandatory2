package main

import (
	"context"
	"log"
	"os"

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
	peers     []grpc.ClientConn
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

//det her er "multicast" delen af algoritmen
func (c *diMutexClient) RequestAccess(ctx context.Context, in *d.AccessRequest, opts ...grpc.CallOption) (*d.AccessGrant, error) {
	//nu skal vi på en måde tale med alle og finde "vinderen"
	peers := c.peers
	for i := 0; i < len(peers); i++ {
		//peers[i].Connect()
		//peers[i].Invoke()
	}
	//compiler gets happy
	return nil, nil
}

//"konverterer" serf-medlemmer til grpc-forbindelser (måske virker det? måske er det for overkompliceret?)
func setupConnection(c *diMutexClient) {
	members := getOtherMembers(c.cluster)
	for i := 0; i < len(members); i++ {
		addr := members[i].Addr.String()
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)

		}
		c.peers[i] = *conn
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
