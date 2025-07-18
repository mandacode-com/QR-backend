package router

import (
	"github.com/gin-gonic/gin"
	"qr.mandacode.com/redirect/ent"
	"qr.mandacode.com/redirect/internal/handler"
)

// RegisterRoutes sets up all routes with the given Ent client
func RegisterRoutes(r *gin.Engine, client *ent.Client) {
	r.GET("/health", handler.HealthCheckHandler(client))
	r.GET("/:id", handler.RedirectHandler(client))
}
