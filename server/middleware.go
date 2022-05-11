package server

import (
	"net/http"
	"strings"
	"try-bank/token"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationKey = "Authorization"
	AuthBearerKey    = "bearer"
	PayloadKey       = "userPayload"
)

func (srv Server) AuthBearer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(AuthorizationKey)

		if len(authHeader) < 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "authorization not provide"})
			return
		}

		authValue := strings.Split(authHeader, " ")
		if len(authValue) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid authorization header"})
			return
		}

		if strings.ToLower(authValue[0]) != AuthBearerKey {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unsupported authorization"})
			return
		}

		tokenProvider, err := token.NewJWT(srv.env.TokenSymetricKey)
		if err == token.ErrSecretKeyLeng {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
			return
		}

		payload, err := tokenProvider.Verify(authValue[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid authorization header"})
			return
		}

		ctx.Set(PayloadKey, payload)
		ctx.Next()
	}
}
