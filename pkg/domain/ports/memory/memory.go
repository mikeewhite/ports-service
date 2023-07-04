package memory

import (
	"sync"

	"github.com/mikeewhite/ports-service/pkg/domain/ports"
)

// Repository is an in-memory implementation of the port repository
type Repository struct {
	ports map[string]ports.Port
	sync.Mutex
}

// New creates a new in-memory repository
func New() *Repository {
	return &Repository{
		ports: make(map[string]ports.Port),
	}
}

// Add will add a new port to the repository
// Existing entries will be overwritten
func (r *Repository) Add(ports []ports.Port) error {
	r.Lock()
	defer r.Unlock()
	for _, p := range ports {
		r.ports[p.ID] = p
	}
	return nil
}

// Get retrieves the port with the given key
func (r *Repository) Get(id string) (*ports.Port, error) {
	if port, ok := r.ports[id]; ok {
		return &port, nil
	}
	return nil, nil
}

// List returns all ports that are currently stored in the in-memory repository
func (r *Repository) List() ([]ports.Port, error) {
	var storedPorts []ports.Port
	for _, p := range r.ports {
		storedPorts = append(storedPorts, p)
	}
	return storedPorts, nil
}
