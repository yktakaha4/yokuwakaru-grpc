package main

import (
	"context"
	"log"
	"os"

	pb "github.com/yktakaha4/yokuwakaru-grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/resolver"
)

func main() {
	resolver.Register(&exampleResolverBuilder{})

	addr := "testScheme:///localhost"
	creds, err := credentials.NewClientTLSFromFile("localhost.crt", "")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds), grpc.WithBalancerName("round_robin"))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := os.Args[1]

	ctx := context.Background()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

type exampleResolverBuilder struct{}

func (*exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &exampleResolver{
		target: target,
		cc:     cc,
		addrStore: map[string][]string{
			"localhost": {"localhost:50051", "localhost:50052"},
		},
	}
	r.start()
	return r, nil
}
func (*exampleResolverBuilder) Scheme() string { return "testScheme" }

type exampleResolver struct {
	target    resolver.Target
	cc        resolver.ClientConn
	addrStore map[string][]string
}

func (r *exampleResolver) start() {
	addrStrs := r.addrStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
func (*exampleResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (*exampleResolver) Close()                                  {}
