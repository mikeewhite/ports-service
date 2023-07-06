package ingest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pb "github.com/mikeewhite/ports-service/generated/go/ports_service/v1"
	"github.com/mikeewhite/ports-service/pkg/config"
)

type MockPortsClient struct {
	ports []*pb.Port
}

func (mpc *MockPortsClient) AddPorts(_ context.Context, ports []*pb.Port) error {
	mpc.ports = append(mpc.ports, ports...)
	return nil
}

func TestIngest(t *testing.T) {
	cfg := config.New()
	cfg.IngestFilepath = "testdata/ports.json"
	mpc := &MockPortsClient{}
	s := NewService(*cfg, mpc)

	require.NoError(t, s.Parse(context.Background()))
	assert.Len(t, mpc.ports, 1632)

	// check the first port is well-formed
	port := mpc.ports[0]
	require.NotNil(t, port)
	assert.Equal(t, "AEAJM", port.Id)
	assert.Equal(t, "Ajman", port.Name)
	assert.Equal(t, "Ajman", port.City)
	assert.Equal(t, "United Arab Emirates", port.Country)
	assert.Empty(t, port.Alias)
	assert.Empty(t, port.Regions)
	assert.EqualValues(t, []float64{55.5136433, 25.4052165}, port.Coordinates)
	assert.Equal(t, "Ajman", port.Province)
	assert.Equal(t, "Asia/Dubai", port.Timezone)
	assert.EqualValues(t, []string{"AEAJM"}, port.Unlocs)
	assert.Equal(t, "52000", port.Code)
}
