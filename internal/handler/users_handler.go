package handler

import (
	"backend/internal/lib"
	"backend/internal/model"
	"backend/internal/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *svc.UserService
}

func (s *UserHandler) Create(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	res, err := s.svc.Create(&model.UserForm{
		Email:    email,
		Password: password,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Success send data",
			Results: res.Email,
		})
	}
}
