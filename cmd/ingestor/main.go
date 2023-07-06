package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v9"

	"github.com/mikeewhite/ports-service/pkg/clog"
	"github.com/mikeewhite/ports-service/pkg/config"
	"github.com/mikeewhite/ports-service/pkg/grpc"
	"github.com/mikeewhite/ports-service/pkg/ingest"
)

func main() {
	defer clog.Info("Ingestor stopped")
	defer clog.Flush()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	cfg := config.New()
	if err := env.Parse(cfg); err != nil {
		panic(fmt.Errorf("failed to parse config struct: %w\n", err))
	}

	grpcClient := grpc.NewClient(*cfg)
	defer grpcClient.Shutdown()
	service := ingest.NewService(*cfg, grpcClient)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		s := <-signals
		clog.Warnf("Signal received (%s). Shutting down\n", s.String())
		cancel()
	}()
	if err := service.Parse(ctx); err != nil {
		clog.Errorf("Error on ingesting file: %s", err.Error())
	}
}
