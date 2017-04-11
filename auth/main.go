package main

import (
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/tallduck/sailfish-backend/helpers"
	"github.com/tallduck/sailfish-backend/protobuf/auth"
)

type authServer struct {
}

func (s *authServer) Authenticate(ctx context.Context, request *auth.Request) (*auth.Response, error) {
	fmt.Println("Starting authentication request")

	if request.Token == "letmein" {
		return &auth.Response{Status: true}, nil
	}

	return &auth.Response{Status: false}, nil
}

func (s *authServer) Invalidate(ctx context.Context, request *auth.Request) (*auth.Response, error) {
	return &auth.Response{Status: false}, nil
}

func main() {
	port := helpers.GetEnv("APP_PORT", "8080")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	auth.RegisterAuthServer(grpcServer, new(authServer))

	fmt.Println(fmt.Printf("Listening on tcp %v", port))
	grpcServer.Serve(listen)
}
