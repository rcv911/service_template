package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	Addr              string
	ReadHeaderTimeout time.Duration
	ShutdownTimeout   time.Duration
}

type HTTPServer struct {
	server *http.Server
	router *gin.Engine
	mux    *runtime.ServeMux
	logger *zerolog.Logger

	shutdownTimeout time.Duration
}

func NewHTTPServer(cfg Config, logger *zerolog.Logger) *HTTPServer {
	router := routes("debug")
	mux := runtime.NewServeMux()

	setupSwaggerRoutes(router)

	return &HTTPServer{
		server: &http.Server{
			Addr:              cfg.Addr,
			Handler:           router,
			ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		},
		router: router,
		mux:    mux,
		logger: logger,

		shutdownTimeout: cfg.ShutdownTimeout,
	}
}

func (s *HTTPServer) Start(ctx context.Context) error {
	errCh := make(chan error)

	go func() {
		errCh <- s.router.Run(s.server.Addr)
		//errCh <- s.server.ListenAndServe() ?
	}()

	go func() {
		<-ctx.Done()
		errCh <- ctx.Err()
	}()

	return <-errCh
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown http server: %w", err)
	}

	return nil
}

func (s *HTTPServer) Router() *gin.Engine {
	return s.router
}

func (s *HTTPServer) Mux() *runtime.ServeMux {
	return s.mux
}

func setupSwaggerRoutes(router *gin.Engine) {
	// Swagger маршруты
	router.StaticFile("/swagger/cat_admin_v1.swagger.json", "./docs/cat_admin_v1/cat_admin_v1.swagger.json")
	router.StaticFile("/swagger/cat_v1.swagger.json", "./docs/cat_v1/cat_v1.swagger.json")

	// http://localhost:8080/swagger/cat_admin_v1/index.html
	router.GET("/swagger/cat_admin_v1/*any", ginswagger.WrapHandler(swaggerfiles.Handler, ginswagger.URL("/swagger/cat_admin_v1.swagger.json")))
	// http://localhost:8080/swagger/cat_v1/index.html
	router.GET("/swagger/cat_v1/*any", ginswagger.WrapHandler(swaggerfiles.Handler, ginswagger.URL("/swagger/cat_v1.swagger.json")))
}

func routes(mode string) *gin.Engine {
	gin.SetMode(mode)

	r := gin.Default()

	r.Use(gin.Recovery())

	r.Handle(http.MethodGet, "/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	r.Handle(http.MethodGet, "/readiness", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	r.Handle(http.MethodGet, "/metrics", func(ctx *gin.Context) {
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})

	return r
}
