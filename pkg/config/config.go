package config

// Config defines the configuration for this application
type Config struct {
	IngestFilepath string `env:"INGEST_FILEPATH"`
	GRPCServer     *GRPCServer
}

// GRPCServer defines the configuration for a gRPC server
type GRPCServer struct {
	ReflectionEnabled bool   `env:"GRPC_REFLECTION_ENABLED"` // Indicates if the server endpoints can be discovered rather than requiring the source protobuf files
	Address           string `env:"GRPC_SERVER_ADDRESS" envDefault:":9000"`
}

// New creates a new Config instance
func New() *Config {
	return &Config{
		GRPCServer: &GRPCServer{},
	}
}
