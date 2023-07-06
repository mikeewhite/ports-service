package grpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/mikeewhite/ports-service/generated/go/ports_service/v1"
	"github.com/mikeewhite/ports-service/pkg/clog"
	"github.com/mikeewhite/ports-service/pkg/config"
)

const (
	connectionTimeout = 10 * time.Second
)

// Client is used to access the ports gRPC interface
type Client struct {
	portsClient pb.PortsServiceClient
	connection  *grpc.ClientConn
}

// NewClient establishes a gRPC connection and creates a new client
func NewClient(cfg config.Config) *Client {
	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()

	addr := cfg.GRPCServer.Address
	grpcConn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(fmt.Errorf("failed to establish gRPC client connection to %s: %w", addr, err))
	}

	return &Client{
		portsClient: pb.NewPortsServiceClient(grpcConn),
		connection:  grpcConn,
	}
}

// AddPorts adds one or more ports
func (c *Client) AddPorts(ctx context.Context, portsToAdd []*pb.Port) error {
	req := &pb.AddPortsRequest{}
	for _, p := range portsToAdd {
		req.Ports = append(req.GetPorts(), p)
	}
	_, err := c.portsClient.AddPorts(ctx, req)
	if err != nil {
		// TODO - return a custom error that indicates whether the error is fatal (e.g. INVALID_ARGUMENT) and should not
		//        be retried, or non-fatal and can be retried
		return fmt.Errorf("gRPC call to AddPorts failed: %w", err)
	}

	return nil
}

// Shutdown will close the gRPC client connection and should be called on application shutdown
func (c *Client) Shutdown() {
	err := c.connection.Close()
	if err != nil {
		clog.Errorf("error on closing gRPC connection: %s", err.Error())
	}
}
