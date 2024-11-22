package grpcserver

import "google.golang.org/grpc"

type Config struct {
	Addr string
}

type GRPCServer struct {
	server *grpc.Server
}

func New(cfg Config) *GRPCServer {
	return &GRPCServer{
		server: grpc.NewServer(),
	}
}

// Встраивание в HTTP/2 сервер: gRPC работает поверх HTTP/2, поэтому сервер можно встроить в HTTP сервер:
//func StartGRPCServerWithHTTP2() {
//	mux := http.NewServeMux()
//
//	grpcServer := grpc.NewServer()
//	// Регистрация gRPC сервисов
//
//	// Встроить gRPC сервер в HTTP сервер
//	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		grpcServer.ServeHTTP(w, r)
//	}))
//
//	server := &http.Server{
//		Addr:    ":8081",
//		Handler: mux,
//	}
//	if err := server.ListenAndServe(); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
