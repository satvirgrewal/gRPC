package main

import (
	"context"
	"fmt"
	"gRPC/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type CalculatorServer struct {
}

func (*CalculatorServer) AddTwoNumbers(ctx context.Context, req *calculatorpb.AdditionRequest) (*calculatorpb.AdditionResponse, error) {
	fmt.Println("Add Two Numbers function invoked")
	a := req.GetFirstNumber()
	b := req.GetSecondNumber()
	sum := a + b
	res := calculatorpb.AdditionResponse{Sum: sum}
	return &res, nil
}

func main() {
	fmt.Println("Starting calculator server")
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		fmt.Printf("Failure in starting calculator start. Error:%v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(s, &CalculatorServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
