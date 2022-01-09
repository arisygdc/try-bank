package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) WebLogout(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    SESCookieName,
		Value:   SESLogout,
		Expires: time.Now().Add(1 * time.Hour),
	})
	ctx.Redirect(http.StatusFound, LocLogin.RequestURI())
}
