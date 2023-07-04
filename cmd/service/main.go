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
)

func main() {
	defer clog.Info("Service stopped")
	defer clog.Flush()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		s := <-signals
		clog.Warnf("\nSignal received (%s). Shutting down\n", s.String())
		cancel()
	}()

	cfg := config.New()
	if err := env.Parse(cfg); err != nil {
		panic(fmt.Errorf("failed to parse config struct: %w\n", err))
	}

	server := grpc.NewServer(cfg.GRPCServer)
	if err := server.Serve(ctx); err != nil {
		clog.Errorf("gRPC server stopped due to error: %s\n", err.Error())
	}
}
