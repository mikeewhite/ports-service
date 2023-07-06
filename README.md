# Ports Service

Contains two applications for managing ports (that have been defined via a JSON file):
* A service which exposes a gRPC interface for persisting and listing ports
* An ingestor which will parse a given JSON file, extract the port information, and call the service to persist the data

## Prerequisites
* `go` 1.20+
* `golangci-lint`
* `buf` (and associated plugins for generating code)
* `docker`
* `make`

## Usage

### Docker
The applications can be started via Docker using:
```bash
docker compose up
```
The gRPC interface will be exposed locally on port `9000` and can be invoked with, for example, [grpcui](https://github.com/fullstorydev/grpcui):
```bash
grpcui -plaintext localhost:9000
```

### Local
The code can be built, linted, and tested using the provided `Makefile`:
```bash
make lint
make test
make build
```
The applications get then be started with:
```bash
export GRPC_SERVER_ADDRESS="localhost:9000"
export GRPC_REFLECTION_ENABLED="true"
./ports-service
./ports-ingestor
```

## TODOs
- [] Move protobuf code generation into Docker (and remove generated files from source control)
