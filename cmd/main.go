package main

import (
	"backend/internal/handler"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool
	Message string
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.Header(
			"X-Data-Backend", "Koda",
		)
		ctx.JSON(200, Response{
			Success: true,
			Message: "Success access backend",
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")
		fmt.Println(email)
		fmt.Println(password)
		ctx.JSON(200, Response{
			Success: true,
			Message: "Success send data",
		})
	})

	r.POST("/login-res", func(ctx *gin.Context) {
		email := ctx.PostForm("email")
		password := ctx.PostForm("passoword")
		if email == "admin@mail.com" && password == "123" {
			ctx.JSON(200, Response{
				Success: true,
				Message: "Success send data",
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, Response{
				Success: false,
				Message: "Salah bos",
			})
		}
	})

	r.POST("/login-em", func(ctx *gin.Context) {
		h := handler.UserHandler{}
		h.Create(ctx)
	})

	r.GET("/hello-world", func(ctx *gin.Context) {
		ctx.Data(200, "text/html", []byte(
			`<h1>Hello World</h1>`,
		))
	})

	r.Run("0.0.0.0:8080")
}
