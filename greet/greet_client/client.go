package main

import (
	"context"
	"fmt"
	"gRPC/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing in client: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do unary RPC")
	fmt.Println("Created client", c)
	greeting := greetpb.Greeting{
		FirstName: "Satvir",
		LastName:  "Grewal",
	}
	req := greetpb.GreetRequest{
		Greeting: &greeting,
	}
	res, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error while calling Greet, %v", err)
	}
	fmt.Printf("Response: %v", res.Result)
}
