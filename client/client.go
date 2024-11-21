package main

import (
	"context"
	"fmt"
	"grpc_demo/greeting"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051",grpc.WithInsecure(),grpc.WithBlock())
	if err != nil{
		log.Fatalf("did not connect: %v",err)
	}
	defer conn.Close()

	c := greeting.NewGreetingServiceClient(conn)

	req := &greeting.GreetingRequest{Name: "john"}
	res,err := c.GetGreeting(context.Background(),req)
	if err != nil{
		log.Fatalf("could not get greeting: %v",err)
	}
	fmt.Println(res.GetMessage())
}