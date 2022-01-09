package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) ViewVA(ctx *gin.Context) {
	session, err := ctx.Cookie(SESCookieName)
	log.Println(session == SESLogout, session == "")
	if err != nil || session == SESLogout || session == "" {
		ctx.Redirect(http.StatusFound, LocLogin.RequestURI())
		return
	}

	regNum, err := strconv.ParseInt(session, 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, LocLogin.RequestURI())
		return
	}
	account, err := ctr.Repo.GetAccount(ctx, int32(regNum))
	if err != nil {
		ctx.Redirect(http.StatusFound, LocLogin.RequestURI())
		return
	}

	user, err := ctr.Repo.GetUser(ctx, account)
	if err != nil {
		ctx.Redirect(http.StatusFound, LocLogin.RequestURI())
		return
	}

	saldo, _ := ctr.Repo.CekSaldo(ctx, int32(regNum))
	ctx.HTML(http.StatusOK, "va.html", gin.H{
		"saldo":     saldo,
		"firstname": user.Firstname,
		"lastname":  user.Lastname,
	})
}
