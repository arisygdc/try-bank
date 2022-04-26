package virtualaccount

import (
	"net/http"
	"try-bank/request"

	"github.com/gin-gonic/gin"
)

func (va VirtualAccountController) Register(ctx *gin.Context) {
	var req request.RegisterVirtualAccount
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

}
