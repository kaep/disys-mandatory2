package main

import (
	"context"
	"fmt"

	d "github.com/kaeppen/disys-mandatory2/dimutex"
	"google.golang.org/grpc"
)

type diMutexClient struct {
	cc        grpc.ClientConnInterface
	state     State
	timestamp int
}

const (
	port = "???"
)

func main() {
	//hvordan skal opsætning af

	c := diMutexClient{}
	//Initialize with state = released
	c.state = Released
	c.timestamp = 0 //init logical clock to 0
	fmt.Println("fuck imports")

}

func (c *diMutexClient) RequestAccess(ctx context.Context, in *d.AccessRequest, opts ...grpc.CallOption) (*d.AccessGrant, error) {
	//kode her :)
	c.state = Wanted //Set the state to wanted

	//compiler gets happy
	return nil, nil
}
