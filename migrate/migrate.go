package main

import (
	"github.com/abdullahalsazib/crud_app/initializers"
	"github.com/abdullahalsazib/crud_app/models"
)

func init() {
	initializers.LoadEnvVarialbe()
	initializers.ConnectTodb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
