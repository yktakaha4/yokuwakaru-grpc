package main

import (
	"context"
	"errors"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	pb "github.com/yktakaha4/yokuwakaru-grpc"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)

	time.Sleep(3 * time.Second)

	if in.Name == "go-error" {
		return nil, errors.New("test error")
	} else if in.Name == "grpc-error" {
		return nil, status.New(codes.NotFound, "resource not found").Err()
	} else if in.Name == "grpc-error-details" {
		st, _ := status.New(codes.Aborted, "aborted").WithDetails(&errdetails.RetryInfo{
			RetryDelay: &duration.Duration{
				Seconds: 3,
				Nanos:   0,
			},
		})
		return nil, st.Err()
	}

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	cred, err := credentials.NewServerTLSFromFile("localhost.crt", "localhost.key")

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.Creds(cred))
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("gRPC server listening on " + addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
