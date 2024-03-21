package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/miladimos/sanjabi/app/routes"
)

type User struct {
	FirstName string `form: "firstname" binding: "required"`
	LastName  string `form: "lastname" binding: "required"`
}

func user(context *gin.Context) {
	name := context.Param("name")
	context.JSON(http.StatusOK, gin.H{"msg": name})
}

func main() {
	// port := os.Getenv("PORT")

	r := routes.SetupRouter()
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
