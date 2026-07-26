package handlers

import (
	"backend/internals/libs"
	"backend/internals/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc services.UserServices
}

func (h *UserHandler) GetAvailUsers(ctx *gin.Context) {
	var query libs.QueryParams
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "Invalid Query Params",
		})
		return
	}

	query.SEARCH_PAR = ctx.QueryMap("search")

	res := h.svc.GetAvailUsers(&query)
	ctx.JSON(http.StatusAccepted, &libs.UserResponseHttp{
		Success: true,
		Status:  http.StatusAccepted,
		Message: "Success get data",
		Queries: &query,
		Results: res,
	})
}
