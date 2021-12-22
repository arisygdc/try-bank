package controller

import (
	"net/http"
	"try-bank/request"

	"github.com/gin-gonic/gin"
)

func (c Controller) CreateUser(ctx *gin.Context) {
	var req request.PostUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.DBSource.CreateUser(ctx, req, "basic"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})
}
