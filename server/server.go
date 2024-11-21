package main

import (
	"context"
	"fmt"
	"grpc_demo/greeting"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	greeting.UnimplementedGreetingServiceServer
}

func (s *server)GetGreeting(ctx context.Context,req *greeting.GreetingRequest)(*greeting.GreetingResponse,error){
	message := fmt.Sprintf("Hello, %s!",req.GetName())
	return &greeting.GreetingResponse{Message: message},nil
}

func main(){
	lis,err := net.Listen("tcp",":50051")
	if err != nil{
		log.Fatalf("failed to listen: %v",err)
	}

	s := grpc.NewServer()

	greeting.RegisterGreetingServiceServer(s,&server{})

	fmt.Println("Server started at :50051")
	if err := s.Serve(lis);err != nil{
		log.Fatalf("failed to serve: %v",err)
	}
}