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
		auth.POST("/login", a.Login)
		auth.POST("/register", a.Register)
		auth.POST("/forgot-password")
	}

	users := r.Group("/users")
	{
		users.GET("", h.GetAvailUsers)
		users.GET("/:id")
	}
	profiles := r.Group("/profiles")
	{
		profiles.GET("")
		profiles.GET("/cart/:id")
		profiles.POST("/cart/:id")
		profiles.PATCH("/cart/:id")
		profiles.GET("/address/:id")
		profiles.POST("/address/:id")
		profiles.GET("/favorites/:id")
	}

	products := r.Group("/products")
	{
		products.GET("")
		products.GET("/:slugs")
	}

	r.Run(":8080")
}
