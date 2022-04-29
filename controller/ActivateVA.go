package controller

import (
	"net/http"
	"try-bank/request"

	"github.com/gin-gonic/gin"
)

func (c DeprecatedController) ActivateVA(ctx *gin.Context) {
	var req request.VirtualAccount
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	vaIdentity, vaKey, err := c.Repo.ActivateVA(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "created",
		"va_number": vaIdentity,
		"va_key":    vaKey,
	})
}
