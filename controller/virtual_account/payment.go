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

	va_identity, va_number, err := ctr.service.VirtualAccount.ValidateVirtualAccount(req.VirtualAccount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	payerAccount, err := ctr.service.Account.CustomerAccount(ctx, req.Payer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ca, err := ctr.service.VirtualAccount.VirtualAccountGetCompany(ctx, int32(va_identity))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	issuedPaymentParam := virtualaccount.IssueVAPayment{
		Virtual_account_id:     ca.VirtualAccountID,
		Virtual_account_number: int32(va_number),
		Payment_charge:         700000,
	}

	issuedPayment, err := ctr.service.VirtualAccount.IssuedVAPaymentValidation(ctx, issuedPaymentParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	paid, err := ctr.service.VirtualAccount.PaymentVirtualAccount(ctx, virtualaccount.PayVA{
		IssuedPayment: issuedPayment.ID,
		PayerWallet:   payerAccount.WalletID,
	})

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
