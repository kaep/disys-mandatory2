package main

import (
	"context"
	"fmt"
	"https://github.com/kaeppen/disys-mandatory2/DiMutex"
	"google.golang.org/grpc"
)


type diMutexClient struct {
	cc grpc.ClientConnInterface
	state State
	timestamp int 
}

const (
	port = "???"
)

func main() {
	//Initialize with state = released
	state = Released
	timestamp = 0 //init logical clock to 0
	fmt.Println("fuck imports")

}

func (c *diMutexClient) RequestAccess(ctx context.Context, in *AccessRequest, opts ...grpc.CallOption) (*AccessGrant, error) {
	//kode her :) 
	c.state = Wanted //Set the state to wanted 
	
}
