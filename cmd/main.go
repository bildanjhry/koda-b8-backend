package main

import (
	"backend/internals/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	h := handlers.UserHandler{}
	a := handlers.AuthHandlers{}

	auth := r.Group("/auth")
	{
		auth.POST("/login")
		auth.POST("/register", a.Login)
		auth.POST("/forgot-password")
	}

	users := r.Group("/users")
	{
		users.GET("", h.GetAvailUsers)
		users.GET("/:id")
	}

	r.Run(":8080")
}
