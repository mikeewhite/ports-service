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
	"github.com/mikeewhite/ports-service/pkg/domain/ports"
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

	portService ports.Service
}

// NewServer will create a gRPC server and start a TCP listener
func NewServer(cfg *config.GRPCServer, portService ports.Service) *Server {
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
	pb.RegisterPortsServiceServer(grpcSrv, &Service{
		portService: portService,
	})

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
	if len(req.Ports) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "At least one port must be provided")
	}

	var portsToStore []ports.Port
	for _, port := range req.Ports {
		if port.Id == "" {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid port: ID is required")
		}
		portsToStore = append(portsToStore, fromGRPCPort(port))
	}

	if err := s.portService.AddPorts(ctx, portsToStore); err != nil {
		return nil, status.Errorf(codes.Internal, "Error on storing ports: %s", err.Error())
	}

	return &pb.AddPortsResponse{}, nil
}

// ListPorts will list all currently stored ports
func (s *Service) ListPorts(ctx context.Context, _ *pb.ListPortsRequest) (*pb.ListPortsResponse, error) {
	storedPorts, err := s.portService.ListPorts(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error on listing ports: %s", err.Error())
	}
	resp := &pb.ListPortsResponse{}
	for _, port := range storedPorts {
		resp.Ports = append(resp.GetPorts(), toGRPCPort(port))
	}
	return resp, nil
}

func fromGRPCPort(p *pb.Port) ports.Port {
	return ports.Port{
		ID:          p.Id,
		Name:        p.Name,
		City:        p.City,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Coordinates: p.Coordinates,
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}
}

func toGRPCPort(p ports.Port) *pb.Port {
	return &pb.Port{
		Id:          p.ID,
		Name:        p.Name,
		City:        p.City,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Coordinates: p.Coordinates,
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}
}
