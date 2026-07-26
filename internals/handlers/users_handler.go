package handlers

import (
	"backend/internals/libs"
	"backend/internals/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc services.UserServices
}

func (h *UserHandler) GetAvailUsers(ctx *gin.Context) {

	lim, _ := strconv.Atoi(ctx.Query("limit"))
	pag, _ := strconv.Atoi(ctx.Query("page"))

	QueryParams := libs.QueryParams{
		SEARCH_PAR: ctx.Query("name"),
		ORDER_BY:   ctx.Query("sort"),
		ORDER_TYPE: ctx.Query("orders"),
		LIMIT:      lim,
		OFFSET:     pag,
	}

	res := h.svc.GetAvailUsers(&QueryParams)
	ctx.JSON(http.StatusAccepted, &libs.UserResponseHttp{
		Success: true,
		Status:  http.StatusAccepted,
		Message: "Success get data",
		Results: res,
	})
}
