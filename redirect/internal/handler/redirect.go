package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"qr.mandacode.com/redirect/ent"
	"qr.mandacode.com/redirect/ent/qrtarget"
)

// RedirectHandler returns a Gin handler for QR redirection
func RedirectHandler(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		idStr := c.Param("id")

		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
			return
		}

		qr, err := client.QrTarget.
			Query().
			Where(qrtarget.IDEQ(id)).
			WithTargettype().
			Only(ctx)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "QR target not found"})
			return
		}

		switch qr.Edges.Targettype.Type {
		case "url":
			ResolveRedirectUrl(c, qr)
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported target type"})
			return
		}
	}
}

// ResolveRedirectUrl handles URL redirection
func ResolveRedirectUrl(c *gin.Context, qr *ent.QrTarget) {
	if qr.Target != nil && *qr.Target != "" {
		c.Redirect(http.StatusFound, *qr.Target)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "target URL is empty"})
}
