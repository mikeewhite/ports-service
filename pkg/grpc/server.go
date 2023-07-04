package grpc

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/mikeewhite/ports-service/generated/go/ports_service/v1"
	"github.com/mikeewhite/ports-service/pkg/clog"
	"github.com/mikeewhite/ports-service/pkg/config"
)

// Server represents a gRPC server
type Server struct {
	address  string
	listener *net.Listener
	server   *grpc.Server
}

// Service defines a gRPC service that will be exposed by a server
type Service struct {
	pb.UnimplementedPortsServiceServer
}

// NewServer will create a gRPC server and start a TCP listener
func NewServer(cfg *config.GRPCServer) *Server {
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		panic(fmt.Errorf("failed to initialise TCP listener: %w", err))
	}

	grpcSrv := grpc.NewServer()
	srv := &Server{
		address:  listener.Addr().String(),
		server:   grpcSrv,
		listener: &listener,
	}

	if cfg.ReflectionEnabled {
		reflection.Register(grpcSrv)
	}

	// register the service
	pb.RegisterPortsServiceServer(grpcSrv, &Service{})

	return srv
}

// Serve will initialise the gRPC server so that it serve traffic
func (s *Server) Serve(ctx context.Context) error {
	clog.Infof("Starting gRPC server at %s\n", s.address)
	go func() {
		<-ctx.Done()
		clog.Info("Stopping gRPC server")
		s.server.GracefulStop()
	}()
	return s.server.Serve(*s.listener)
}

// AddPorts adds one or more ports
func (s *Service) AddPorts(ctx context.Context, req *pb.AddPortsRequest) (*pb.AddPortsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "TODO")
}

// ListPorts will list all currently stored ports
func (s *Service) ListPorts(ctx context.Context, req *pb.ListPortsRequest) (*pb.ListPortsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "TODO")
}
