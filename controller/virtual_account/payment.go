package virtualaccount

import (
	"net/http"
	virtualaccount "try-bank/app_service/virtual_account"

	"github.com/gin-gonic/gin"
)

// Payer is customer (registered_number) wants to pay vitrual account
type PostVirtualAccountPay struct {
	Payer          int32  `json:"registered_number" binding:"required"`
	VirtualAccount string `json:"virtual_account" binding:"required"`
}

// TODO
// payer should be in authentication barier
// http status code
func (ctr VirtualAccountController) VirtualAccount_pay(ctx *gin.Context) {
	var req PostVirtualAccountPay
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	vaID, err := ctr.service.VirtualAccount.VirtualAccountID(ctx, 3)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	issuedPayment, err := ctr.service.VirtualAccount.CheckIssuedVAPayment(ctx, virtualaccount.IssueVAPayment{
		Virtual_account_id:     vaID,
		Virtual_account_number: req.Payer,
		Payment_charge:         700000,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	paid, err := ctr.service.VirtualAccount.PaymentVirtualAccount(ctx, issuedPayment.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"amount":    issuedPayment.Payment_charge,
		"issued_at": issuedPayment.IssuedAt,
		"paid_at":   paid.PaidAt,
	})
}
