package company

import (
	"net/http"
	"try-bank/app_service/company"

	"github.com/gin-gonic/gin"
)

type PostCompany struct {
	Name    string  `json:"name" binding:"required"`
	Email   string  `json:"email" binding:"required"`
	Phone   string  `json:"phone" binding:"required"`
	Pin     string  `json:"pin" binding:"required"`
	Deposit float64 `json:"deposit" binding:"required"`
}

// TODO
// validation
func (ctr CompanyController) RegisterCompany(ctx *gin.Context) {
	var req PostCompany
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	registered, err := ctr.service.Company.CreateCompanyAccount(ctx, company.RegisterCompanyParam{
		PublicInfo_company: company.PublicInfo_company{
			Name:  req.Name,
			Email: req.Email,
			Phone: req.Phone,
		},
		Pin:     req.Pin,
		Deposit: req.Deposit,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"name":              registered.Name,
		"email":             registered.Email,
		"phone":             registered.Phone,
		"deposit":           registered.Deposit,
		"registered_number": registered.RegisterNumber,
	})
}
