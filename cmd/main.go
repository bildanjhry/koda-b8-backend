package main

import (
	"backend/internals/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	h := handlers.UserHandler{}

	auth := r.Group("/users")
	{
		auth.GET("", h.GetAvailUsers)
		auth.GET("/:id")
	}

	r.Run(":8080")
}
