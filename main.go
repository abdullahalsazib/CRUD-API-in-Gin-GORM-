package main

import (
	"github.com/abdullahalsazib/crud_app/controllers"
	"github.com/abdullahalsazib/crud_app/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVarialbe()
	initializers.ConnectTodb()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.Run() // listen and serve on 0.0.0.0:8080
}
