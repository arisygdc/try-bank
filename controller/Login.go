package controller

import (
	"net/http"
	"net/url"
	"strconv"
	"time"
	"try-bank/request"

	"github.com/gin-gonic/gin"
)

var (
	SESLogout     = "LOGOUT"
	SESCookieName = "X-SESSION"
	LocVA         = url.URL{Path: "/payment/va", RawQuery: url.Values{}.Encode()}
	LocLogin      = url.URL{Path: "/login", RawQuery: url.Values{}.Encode()}
)

func (ctr Controller) Login(ctx *gin.Context) {
	regNumParam := ctx.PostForm("register_number")
	pinParam := ctx.PostForm("pin")
	regnum, err := strconv.ParseInt(regNumParam, 10, 32)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "index.html", gin.H{})
		return
	}
	var req = request.LoginReq{
		RegNum: int32(regnum),
		Pin:    pinParam,
	}

	_, err = ctr.Repo.Login(ctx, req.RegNum, req.Pin)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "index.html", gin.H{})
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    SESCookieName,
		Value:   regNumParam,
		Expires: time.Now().Add(1 * time.Hour),
	})

	ctx.Redirect(http.StatusFound, LocVA.RequestURI())
}
