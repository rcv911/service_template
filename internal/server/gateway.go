package server

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rcv911/service_template/pkg/cat_admin_v1"
	"github.com/rcv911/service_template/pkg/cat_v1"
	"google.golang.org/grpc"
)

func RegisterGRPCHandlers(ctx context.Context, router *gin.Engine, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	if err := cat_admin_v1.RegisterCatAdminServiceHandler(ctx, mux, conn); err != nil {
		return fmt.Errorf("could not register cat admin v1 gRPC gateway: %v", err)
	}

	if err := cat_v1.RegisterCatServiceHandler(ctx, mux, conn); err != nil {
		return fmt.Errorf("could not register cat v1 gRPC gateway: %v", err)
	}

	router.GET("/v1/cats", gin.WrapH(mux))        // GET /v1/cats
	router.POST("/v1/cats", gin.WrapH(mux))       // POST /v1/cats
	router.GET("/v1/cats/:id", gin.WrapH(mux))    // GET /v1/cats/{id}
	router.PUT("/v1/cats/:id", gin.WrapH(mux))    // PUT /v1/cats/{id}
	router.DELETE("/v1/cats/:id", gin.WrapH(mux)) // DELETE /v1/cats/{id}

	return nil
}
