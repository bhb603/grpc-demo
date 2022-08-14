package calculator

import (
	"context"
	"log"

	"github.com/bhb603/grpc-demo/golang/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Calculator implements the grpc CalculatorServer.
type Calculator struct {
	pb.UnimplementedCalculatorServer
}

// NthFibonacci implements calculator.NthFibonacci
func (c *Calculator) NthFibonacci(ctx context.Context, params *pb.FibonacciParams) (*pb.NthFibonacciResponse, error) {
	var result int32

	n := params.GetN()
	log.Printf("Received: request for %dth fibonacci", n)

	if n < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "cannot get a negative fibonacci")
	}

	if n > 46 {
		return nil, status.Errorf(codes.InvalidArgument, "too large for 32-bit integers")
	}

	if n <= int32(1) {
		result = n
		return &pb.NthFibonacciResponse{Result: result}, nil
	}

	var a, b int32 = 0, 1
	for i := 2; i <= int(n); i++ {
		result = a + b
		a = b
		b = result
	}

	return &pb.NthFibonacciResponse{Result: result}, nil
}

// Sum implements calculator.Sum
func (c *Calculator) Sum(ctx context.Context, params *pb.SumParams) (*pb.SumResponse, error) {
	log.Printf("Received Sum request: numbers=%v", params.GetNumbers())

	var sum int32
	for _, n := range params.Numbers {
		sum += n
	}

	return &pb.SumResponse{Sum: sum}, nil
}
