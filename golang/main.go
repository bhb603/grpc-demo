package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/bhb603/grpc-demo/golang/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// Calculator implements the CalculatorServer.
type Calculator struct {
	pb.UnimplementedCalculatorServer
}

// NthFibonacci implements calculator.NthFibonacci
func (c *Calculator) NthFibonacci(ctx context.Context, in *pb.FibonacciParams) (*pb.NthFibonacciResponse, error) {
	var result int32

	n := in.GetN()
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

// Substrings implements calculator.Subsctrings
func (c *Calculator) Substrings(ctx context.Context, in *pb.SubstringsParams) (*pb.SubstringsResponse, error) {
	log.Printf("Received request for substrings of %q", in.GetStr())

	chars := []rune(in.GetStr())
	substrings := make(map[string]struct{})

	for i := 0; i < len(chars); i++ {
		for j := i + 1; j < len(chars); j++ {
			substr := string(chars[i:j])
			substrings[substr] = struct{}{}
		}
	}

	list := []string{}
	for key := range substrings {
		list = append(list, key)
	}
	return &pb.SubstringsResponse{Substrings: list}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterCalculatorServer(s, &Calculator{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
