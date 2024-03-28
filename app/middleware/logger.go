package middleware

import (
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		// before request
		context.Next()
		// after request
	}
}
