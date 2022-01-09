package controller

import (
	"net/http"
	"strconv"
	"try-bank/request"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) WebVA(ctx *gin.Context) {
	var req = request.PaymentVA{}

	session, err := ctx.Cookie(SESCookieName)
	if err == nil || session == SESLogout && session == "" {
		ctx.Redirect(http.StatusFound, LocLogin.RequestURI())
		return
	}

	regnum, err := strconv.ParseInt(session, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, LocLogin.RequestURI())
		return
	}

	req.RegNum = int32(regnum)
	req.Pin = ctx.PostForm("pin")
	req.VirtualAccount = ctx.PostForm("va")

	if err := ctr.Repo.PaymentVA(ctx, req); err != nil {
		ctx.Redirect(http.StatusFound, LocVA.RequestURI())
		return
	}

	ctx.Redirect(http.StatusFound, LocVA.RequestURI())
}
