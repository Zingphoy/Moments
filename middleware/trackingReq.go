package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var trackingHeader = "Tracking-Id"

func getTrackingId() string {
	return uuid.New().String()
}

func TrackingId() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tidNew string
		tid := c.GetHeader(trackingHeader)
		if tid == "" {
			tidNew = getTrackingId()
			c.Header(trackingHeader, tidNew)
		}
		c.Set(trackingHeader, tidNew)
		c.Next()
	}
}
