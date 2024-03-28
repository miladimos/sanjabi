package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miladimos/sanjabi/app/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.NoRoute(middleware.NoRouteHandler())
	r.HandleMethodNotAllowed = true
	r.NoMethod(middleware.NoMethodHandler())
	// r.Use(Logger())

	api := r.Group("/api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("/users", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"users": "users"})
			})

			v1.GET("/users/:user", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"user": "user"})
			})
		}
	}

	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	// r.GET("/users/:name", user)

	return r
}
