package handlers

import (
	"backend/internals/models"
	"backend/internals/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandlers struct {
	svc *services.AuthServices
}

func (h *AuthHandlers) Login(ctx *gin.Context) {
	var form models.LoginRequest

	errForm := ctx.ShouldBind(&form)
	if errForm != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": errForm.Error(),
		})
		return
	}
	res, errRes := h.svc.Login(&form)
	if errRes != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": errRes.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Success Login",
		"Results": &models.LoginResponse{
			Id:    res.Id,
			Token: res.Token,
		},
	})
}

func (h *AuthHandlers) Register(ctx *gin.Context) {
	var form models.RegisterRequest

	errForm := ctx.ShouldBind(&form)
	if errForm != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": errForm.Error(),
		})
		return
	}
	res, errRes := h.svc.Register(&form)
	if errRes != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": errRes.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Success": true,
		"Message": "Success Create Account",
		"Result": &models.RegisterResponse{
			Id: res.Id,
		},
	})
}
