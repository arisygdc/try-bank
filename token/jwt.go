package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type JWToken struct {
	secretKey string
}

const minSecretKey = 32

var signingMethod = jwt.SigningMethodHS256

func NewJWT(secretKey string) (*JWToken, error) {
	if len(secretKey) < minSecretKey {
		return nil, fmt.Errorf("invalid secret key length: must be at least %d", minSecretKey)
	}

	return &JWToken{
		secretKey: secretKey,
	}, nil
}

// generating payload: NewPayload()
func (t JWToken) Make(payload *Payload) (string, error) {
	claims := jwt.NewWithClaims(signingMethod, payload)
	generatedToken, err := claims.SignedString([]byte(t.secretKey))
	return generatedToken, err
}

func (t JWToken) Verify(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != signingMethod.Alg() {
			return nil, jwt.ErrTokenUnverifiable
		}

		return []byte(t.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc, jwt.WithoutClaimsValidation())
	if err != nil {
		return nil, err
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil
}
