package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var TrackingHeader = "Tracking-Id"

func getTrackingId() string {
	return uuid.New().String()
}

func TrackingId() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tidNew string
		tid := c.GetHeader(TrackingHeader)
		if tid == "" {
			tidNew = getTrackingId()
			c.Header(TrackingHeader, tidNew)
		}
		c.Set(TrackingHeader, tidNew)
		c.Next()
	}
}
