package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureCors() gin.HandlerFunc {
	return cors.Default()
}
