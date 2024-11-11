package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rcv911/service_template/pkg/cat_admin_v1"
	"google.golang.org/grpc"
)

func httpServerStart(stopCh chan os.Signal) {
	grpcServer := grpc.NewServer()
	cat_admin_v1.RegisterCatAdminServiceServer(grpcServer, &CatAdminServiceServer{})

	// Настройка gRPC-Gateway
	gwMux := runtime.NewServeMux()
	err := cat_admin_v1.RegisterCatAdminServiceHandlerServer(context.Background(), gwMux, &CatAdminServiceServer{})
	if err != nil {
		log.Fatalf("failed to register gateway server: %v", err)
	}

	httpSrv := &http.Server{
		Addr:    ":8080",
		Handler: gwMux,
	}

	// HTTP сервер для REST запросов
	go func() {
		log.Println("Starting HTTP server (gRPC-Gateway) on :8080")
		if err = httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to serve HTTP server: %v", err)
		}
	}()

	<-stopCh
	log.Println("Shutting down HTTP server...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = httpSrv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("failed to gracefully stop HTTP server: %v", err)
	}
	log.Println("HTTP server gracefully stopped")
}
