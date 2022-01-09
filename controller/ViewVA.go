package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) ViewVA(ctx *gin.Context) {
	session, err := ctx.Cookie(SESCookieName)
	log.Println(session == SESLogout, session == "")
	if err != nil || session == SESLogout || session == "" {
		ctx.Redirect(http.StatusFound, LocLogin.RequestURI())
		return
	}
	ctx.HTML(http.StatusOK, "va.html", gin.H{})
}
