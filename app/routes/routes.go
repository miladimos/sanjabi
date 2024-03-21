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
		}
	}

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "Hello"})
	})

	// r.GET("/users/:name", user)

	return r
}
