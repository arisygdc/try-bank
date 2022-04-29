package controller

import (
	"net/http"
	"try-bank/request"

	"github.com/gin-gonic/gin"
)

func (c DeprecatedController) Transfer(ctx *gin.Context) {
	var req request.Transfer
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.Repo.Transfer(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "transfered",
	})
}
