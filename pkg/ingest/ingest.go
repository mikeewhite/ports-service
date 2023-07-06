package ingest

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	pb "github.com/mikeewhite/ports-service/generated/go/ports_service/v1"
	"github.com/mikeewhite/ports-service/pkg/clog"
	"github.com/mikeewhite/ports-service/pkg/config"
)

const (
	batchSize = 10
)

// Port models the JSON representation of a port
type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

// PortsClient defines the methods required to manage ports
type PortsClient interface {
	AddPorts(ctx context.Context, portsToAdd []*pb.Port) error
}

// Service defines the ingest service
type Service struct {
	portsClient PortsClient
	filepath    string
	batch       []*pb.Port
	sync.Mutex
}

// NewService creates a new ingest service
func NewService(cfg config.Config, portsClient PortsClient) *Service {
	return &Service{
		portsClient: portsClient,
		filepath:    cfg.IngestFilepath,
	}
}

// Parse will parse the JSON file at the configured file path, extract the contained port information
// and attempt to call the port client to store the information
func (s *Service) Parse(ctx context.Context) error {
	s.Lock()
	defer s.Unlock()

	file, err := os.Open(s.filepath)
	if err != nil {
		return fmt.Errorf("failed to read ingest filepath '%s': %w", s.filepath, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			clog.Errorf("error on closing file: %s", err.Error())
		}
	}(file)

	dec := json.NewDecoder(file)

	// read opening JSON delimiter
	token, err := dec.Token()
	if err != nil {
		return fmt.Errorf("failed to read opening delimiter: %w", err)
	}
	if token != json.Delim('{') {
		return fmt.Errorf("unexpected token encountered on reading opening delimiter: '%s'", token)
	}

	// read the remaining JSON
	for dec.More() {
		// get key
		portKeyT, err := dec.Token()
		if err != nil {
			return fmt.Errorf("failed to read port key: %w", err)
		}

		// check key is a string
		key, ok := portKeyT.(string)
		if !ok {
			return fmt.Errorf("unexpected type for port key: '%T'", token)
		}

		// read the rest of the port JSON
		var port Port
		err = dec.Decode(&port)
		if err != nil {
			return fmt.Errorf("error on decoding port with key '%s': %w", key, err)
		}

		// add the port to the batch of ports to store
		clog.Infof("Adding port with key '%s' to batch", key)
		s.batch = append(s.batch, toGRPCPort(key, port))
		if err = s.flushBatchIfRequired(ctx); err != nil {
			return err
		}
	}
	if err = s.flushBatch(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Service) flushBatchIfRequired(ctx context.Context) error {
	if len(s.batch) >= batchSize {
		return s.flushBatch(ctx)
	}
	return nil
}

func (s *Service) flushBatch(ctx context.Context) error {
	if len(s.batch) > 0 {
		clog.Infof("Flushing batch of size %d", len(s.batch))
		err := s.portsClient.AddPorts(ctx, s.batch)
		if err != nil {
			// TODO - retry call (if error code is not INVALID_ARGUMENT) a configurable number of times with
			//        an exponential backoff before failing
			return fmt.Errorf("failed to store ports via gRPC interface: %w", err)
		}
	}
	s.batch = []*pb.Port{}
	return nil
}

func toGRPCPort(key string, p Port) *pb.Port {
	return &pb.Port{
		Id:          key,
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
