package sig

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
)

var ErrShutdownSignalReceived = errors.New("shutdown")

func Listen(ctx context.Context) error {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-shutdown:
		return nil
		//return ErrShutdownSignalReceived
	case <-ctx.Done():
		return nil
	}
}
