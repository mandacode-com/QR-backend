package handler

import (
	"context"
	"net/http"
	"qr.mandacode.com/redirect/ent"

	"github.com/gin-gonic/gin"
)

// RedirectHandler returns a Gin handler for QR redirection
func HealthCheckHandler(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		// Perform a simple health check by querying the database
		_, err := client.QrTarget.Query().First(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "service unavailable"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	}
}
