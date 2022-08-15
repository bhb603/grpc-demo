package calculator

import (
	"context"
	"log"
	"math/rand"
	"time"

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
	var value int32

	n := params.GetN()
	log.Printf("Received: request for %dth fibonacci", n)

	if n < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "cannot get a negative fibonacci")
	}

	if n > 46 {
		return nil, status.Errorf(codes.InvalidArgument, "too large for 32-bit integers")
	}

	if n <= int32(1) {
		value = n
		return &pb.NthFibonacciResponse{Value: value}, nil
	}

	var a, b int32 = 0, 1
	for i := 2; i <= int(n); i++ {
		value = a + b
		a = b
		b = value
	}

	return &pb.NthFibonacciResponse{Value: value}, nil
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

func (c *Calculator) RandomStream(params *pb.RandomStreamParams, stream pb.Calculator_RandomStreamServer) error {
	min, max, count := params.GetMin(), params.GetMax(), params.GetCount()
	log.Printf("Received random stream request %d numbers in [%d, %d)", count, min, max)
	if min >= max {
		return status.Errorf(codes.InvalidArgument, "min must be < max")
	}
	for i := int32(0); i < count; i++ {
		n := rand.Int31n(max-min) + min
		if err := stream.Send(&pb.RandomNumber{Value: n}); err != nil {
			log.Printf("done streaming: %v", err)
			return err
		}
		time.Sleep(500 * time.Millisecond)
	}

	log.Printf("done streaming")

	return nil
}
func (c *Calculator) IsPrime(ctx context.Context, params *pb.IsPrimeParams) (*pb.IsPrimeResponse, error) {
	val := params.GetValue()
	log.Printf("Received IsPrime request for %d", val)

	if val <= int32(1) {
		return &pb.IsPrimeResponse{Prime: false}, nil
	}

	for n := int32(2); n*n <= val; n++ {
		if val%n == 0 {
			return &pb.IsPrimeResponse{Prime: false}, nil
		}
	}

	return &pb.IsPrimeResponse{Prime: true}, nil
}
