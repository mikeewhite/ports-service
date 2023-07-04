package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mikeewhite/ports-service/pkg/clog"
)

func main() {
	defer clog.Info("Ingestor stopped")
	defer clog.Flush()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	clog.Info("Ingestor started")

	// Wait until app is closed
	s := <-signals
	clog.Warnf("\nSignal received (%s). Shutting down\n", s.String())
}
