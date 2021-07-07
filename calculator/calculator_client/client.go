package main

import (
	"context"
	"fmt"
	"gRPC/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calling client")
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error dialing client, %v", err)
	}
	defer cc.Close()

	client := calculatorpb.NewCalculatorClient(cc)
	doUnary(client)
}
func doUnary(client calculatorpb.CalculatorClient) {

	addRequest := calculatorpb.AdditionRequest{FirstNumber: 3,
		SecondNumber: 10}
	fmt.Printf("Adding 2 numbers, request: %v", addRequest)
	res, _ := client.AddTwoNumbers(context.Background(), &addRequest)
	fmt.Printf("Sum of numbers, %v", res)
}
