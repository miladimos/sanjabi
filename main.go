package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		c.Next()
		// after request
	}
}

type User struct {
	FirstName string `form: "firstname" binding: "required"`
	LastName  string `form: "lastname" binding: "required"`
}

func user(context *gin.Context) {
	name := context.Param("name")
	context.JSON(http.StatusOK, gin.H{"msg": name})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(Logger())

	api := r.Group("/api")
	{
		api.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"users": "users"})
		})
	}

	r.GET('/', func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "Hello"})
	})

	r.GET("/users/:name", user)

	return r
}

func main() {
	r := setupRouter()
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
