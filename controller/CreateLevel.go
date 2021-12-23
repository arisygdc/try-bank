package controller

import (
	"net/http"
	"try-bank/database/postgres"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionReq struct {
	Name string `from:"name" json:"name" binding:"required"`
}

func (c Controller) CreateLevel(ctx *gin.Context) {
	var req PermissionReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.DBSource.Queries.CreateLevel(ctx, postgres.CreateLevelParams{
		ID:   uuid.New(),
		Name: req.Name,
	}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})
}
