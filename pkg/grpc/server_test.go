package grpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/mikeewhite/ports-service/generated/go/ports_service/v1"
)

func TestAddPorts_InvalidRequest(t *testing.T) {
	s := &Service{}

	tt := []struct {
		name           string
		request        *pb.AddPortsRequest
		expectedErrMsg string
	}{
		{
			name:           "empty request",
			request:        &pb.AddPortsRequest{},
			expectedErrMsg: "At least one port must be provided",
		},
		{
			name:           "invalid port",
			request:        &pb.AddPortsRequest{Ports: []*pb.Port{{Name: "Invalid Port"}}},
			expectedErrMsg: "Invalid port: ID is required",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.request
			resp, err := s.AddPorts(context.Background(), req)
			assert.Nil(t, resp, "expected response to be nil")
			require.NotNil(t, err)

			actErr, ok := status.FromError(err)
			require.True(t, ok)

			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, tc.expectedErrMsg, actErr.Message())
		})
	}
}
