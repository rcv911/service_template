package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/rcv911/service_template/internal/repository"
	"github.com/rcv911/service_template/internal/server"
	"github.com/rcv911/service_template/pkg/pyroscope"
	"github.com/rcv911/service_template/pkg/sig"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	httpHost := "localhost:8080"
	grpcHost := "localhost:8081"

	logger, err := initLogger("info")
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}

	lis, err := net.Listen("tcp", grpcHost)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	client, err := grpc.NewClient(grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed dial gRPC client connection: %v", err)
	}

	repo := repository.New(logger)

	catService := server.NewCatServiceServer(repo)
	catAdminService := server.NewCatAdminServiceServer(repo)

	grpcServer := server.NewGRPCSServer(catService, catAdminService)

	httpServerCfg := server.Config{
		Addr:              httpHost,
		ReadHeaderTimeout: time.Second,
		ShutdownTimeout:   10 * time.Second,
	}

	httpServer := server.NewHTTPServer(httpServerCfg, &zerolog.Logger{})

	err = server.RegisterGRPCHandlers(ctx, httpServer.Router(), httpServer.Mux(), client)
	if err != nil {
		log.Fatalf("failed to register gRPC server handlers: %v", err)
	}

	err = pyroscope.New("service_template", "localhost:4040")
	if err != nil {
		log.Fatalf("failed pyroscope: %v", err)
	}

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return httpServer.Start(gCtx)
	})

	g.Go(func() error {
		return grpcServer.Serve(lis)
	})

	g.Go(func() error {
		return sig.Listen(ctx)
	})

	if err = g.Wait(); err != nil {
		log.Fatalf("failed errgroup: %v", err)
	}

	log.Println("graceful shutdown")

	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()
	log.Println("gRPC server gracefully stopped")

	if err = httpServer.Stop(ctx); err != nil {
		log.Fatalf("failed to shutdown http server: %v", err)
	}
}

func initLogger(level string) (zerolog.Logger, error) {
	l, err := zerolog.ParseLevel(level)
	if err != nil {
		return zerolog.Logger{}, fmt.Errorf("getting log level error [%w]", err)
	}

	return zerolog.New(os.Stdout).Level(l).With().Timestamp().Logger(), nil
}
