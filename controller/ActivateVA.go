package controller

import (
	"net/http"
	"try-bank/request"

	"github.com/gin-gonic/gin"
)

func (c Controller) ActivateVA(ctx *gin.Context) {
	var req request.VirtualAccount
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.DBSource.ActivateVA(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "created",
	})
}
