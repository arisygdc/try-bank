package account

import (
	"net/http"
	"try-bank/app_service/account"
	"try-bank/util"

	"github.com/gin-gonic/gin"
)

type PostCustomerAccount struct {
	Firstname   string  `json:"firstname" binding:"required"`
	Lastname    string  `json:"lastname" binding:"required"`
	Phone       string  `json:"phone" binding:"required"`
	Pin         string  `json:"pin" binding:"required"`
	Birth       string  `json:"birth" binding:"required"`
	TopUp       float64 `json:"topup" binding:"required"`
	AccountType string  `json:"account_type" binding:"required"`
}

// TODO
// validation
func (ctr AccountController) Register(ctx *gin.Context) {
	var req PostCustomerAccount
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	birth, err := util.StrToTime(req.Birth)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "parsing date of birth"})
		return
	}

	account_type, err := ctr.service.Account.GetAccountType(ctx, account.LevelSilver)
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account type plan"})
	if err != nil {
		return
	}

	registerParam := account.CreateCostumerParam{
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Phone:       req.Phone,
		Pin:         req.Pin,
		Birth:       birth,
		TopUp:       req.TopUp,
		AccountType: account_type.ID,
	}

	registeredAccount, err := ctr.service.Account.CreateCustomerAccount(ctx, registerParam)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"name":              registeredAccount.Name,
		"email":             registeredAccount.Email,
		"phone":             registeredAccount.Phone,
		"birth":             registeredAccount.Birth,
		"top_up":            registeredAccount.TopUp,
		"registered_number": registeredAccount.RegisteredNumber,
	})
}
