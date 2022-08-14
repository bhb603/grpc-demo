package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/bhb603/grpc-demo/golang/calculator"
	"github.com/bhb603/grpc-demo/golang/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serverOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(authorizeUnary),
		grpc.ChainStreamInterceptor(authorizeStream),
	}

	s := grpc.NewServer(serverOpts...)
	reflection.Register(s)
	pb.RegisterCalculatorServer(s, &calculator.Calculator{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func authorizeUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("authorizing unary: %s", info.FullMethod)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}

	if !validateApiKey(md) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid api key")
	}

	return handler(ctx, req)
}

func authorizeStream(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("authorizing stream: %s", info.FullMethod)
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return status.Errorf(codes.InvalidArgument, "missing metadata")
	}

	if !validateApiKey(md) {
		return status.Errorf(codes.InvalidArgument, "invalid api key")
	}

	return handler(srv, ss)
}

func validateApiKey(md metadata.MD) bool {
	apiKey := md.Get("x-api-key")
	if len(apiKey) == 0 {
		return false
	}
	return apiKey[0] == "secret"
}
