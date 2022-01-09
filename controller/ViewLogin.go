package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) ViewLogin(ctx *gin.Context) {
	session, _ := ctx.Cookie(SESCookieName)
	session = strings.Trim(session, " ")
	if session != SESLogout {
		if session != "" {
			ctx.Redirect(http.StatusFound, LocVA.RequestURI())
			return
		}
	}
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
