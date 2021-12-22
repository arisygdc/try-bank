package controller

import (
	"net/http"
	"strings"
	"time"
	"try-bank/database/postgres"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserReq struct {
	Firstname string  `from:"firstname" json:"firstname" binding:"required"`
	Lastname  string  `from:"lastname" json:"lastname" binding:"required"`
	Email     string  `from:"email" json:"email" binding:"required"`
	Birth     string  `from:"birth" json:"birth" binding:"required"`
	TopUp     float64 `from:"deposit" json:"deposit" binding:"required"`
	Phone     string  `from:"phone" json:"phone" binding:"required"`
	Pin       string  `from:"pin" json:"pin" binding:"required"`
}

func (c Controller) CreateUser(ctx *gin.Context) {
	var req CreateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	t, err := time.Parse("2006-1-2", strings.Trim(req.Birth, " "))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := postgres.CreateUserParams{
		ID:        uuid.New(),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Birth:     t.UTC(),
		Phone:     req.Phone,
	}

	wallet := postgres.CreateCoustomerWalletParams{
		ID:      uuid.New(),
		Balance: req.TopUp,
	}

	authInfo := postgres.CreateAuthInfoParams{
		ID:               uuid.New(),
		RegisteredNumber: int32(t.Month()) + int32(req.Phone[9]+req.Phone[10]+req.Phone[11]),
	}

	err = c.DBSource.CreateUser(ctx, user, authInfo, wallet, "user")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})
}
