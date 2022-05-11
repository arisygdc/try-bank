package account

import (
	"database/sql"
	"net/http"
	"try-bank/token"

	"github.com/gin-gonic/gin"
)

type PostTransfer struct {
	To      int32   `json:"transfer_to"`
	Balance float64 `json:"balance"`
}

func (ctr AccountController) Transfer(ctx *gin.Context) {
	var req PostTransfer
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	// middleware not yet created
	getPayload, exists := ctx.Get("userPayload")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}

	payload, ok := getPayload.(*token.Payload)
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}

	transferTo, err := ctr.service.Account.CustomerAccount(ctx, req.To)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "registration number not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	transferFrom, err := ctr.service.Account.CustomerAccount(ctx, payload.Registered_number)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "illegal access"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	err = ctr.service.Account.Transfer(ctx, transferFrom.WalletID, transferTo.WalletID, req.Balance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"from":  payload.Registered_number,
		"to":    req.To,
		"total": req.Balance,
	})
}
