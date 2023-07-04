package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mikeewhite/ports-service/pkg/domain/ports"
)

func TestAdd(t *testing.T) {
	repo := New()
	require.NoError(t, repo.Add([]ports.Port{getTestPort()}))

	returnedPorts, err := repo.List()
	require.NoError(t, err)
	require.NotNil(t, returnedPorts)
	require.NotEmpty(t, returnedPorts)

	assert.Len(t, returnedPorts, 1)
	assert.Equal(t, getTestPort(), returnedPorts[0])
}

func TestAdd_ReplacesExisting(t *testing.T) {
	repo := New()
	port := getTestPort()
	require.NoError(t, repo.Add([]ports.Port{port}))

	// mutate the port and store it again
	port.Country = "United Kingdom"
	require.NoError(t, repo.Add([]ports.Port{port}))

	returnedPort, err := repo.Get("AEAJM")
	require.NoError(t, err)
	require.NotNil(t, returnedPort)
	assert.Equal(t, "United Kingdom", port.Country)
}

func getTestPort() ports.Port {
	return ports.Port{
		ID:          "AEAJM",
		Name:        "Ajman",
		City:        "Ajman",
		Country:     "United Arab Emirates",
		Alias:       nil,
		Regions:     nil,
		Coordinates: []float64{55.5136433, 25.4052165},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAJM"},
		Code:        "52000",
	}
}
