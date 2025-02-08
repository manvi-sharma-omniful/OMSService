package router

import (
	"awesomeProject/Project/OMS/controllers"
	"context"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
)

func InternalRoutes(ctx context.Context, server *http.Server) error {
	Order1 := server.Engine.Group("/api/v1", log.RequestLogMiddleware(log.MiddlewareOptions{
		Format:      config.GetString(ctx, "log.format"),
		Level:       config.GetString(ctx, "log.level"),
		LogRequest:  config.GetBool(ctx, "log.request"),
		LogResponse: config.GetBool(ctx, "log.response"),
	}))

	Order1.POST("", controllers.CreateOrder)
	return nil
}
