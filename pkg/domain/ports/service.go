package ports

import (
	"context"
)

// Option is a functional option for configuring the service
type Option func(*Service) error

// Service defines a port service
type Service struct {
	ports PortRepository
}

// NewService creates a nre port service instance
func NewService(opts ...Option) (*Service, error) {
	ps := &Service{}

	for _, option := range opts {
		err := option(ps)
		if err != nil {
			return nil, err
		}
	}
	return ps, nil
}

// WithInMemoryRepository uses the provided in-memory backing repository
func WithInMemoryRepository(repo PortRepository) Option {
	return func(ps *Service) error {
		ps.ports = repo
		return nil
	}
}

// AddPorts stores one or more ports
func (p *Service) AddPorts(_ context.Context, ports []Port) error {
	return p.ports.Add(ports)
}

// ListPorts lists all currently stored ports
func (p *Service) ListPorts(_ context.Context) ([]Port, error) {
	return p.ports.List()
}
