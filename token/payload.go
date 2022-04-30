package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Payload struct {
	SessionID         uuid.UUID `json:"session_id"`
	Registered_number int32     `json:"registered_numer"`
	Issued_at         time.Time `json:"issued_at"`
	Expired_at        time.Time `json:"expired_at"`
	jwt.StandardClaims
}

func Newpayload(RegisteredNumber int32, tokenDuration time.Duration) *Payload {
	timeNow := time.Now()
	return &Payload{
		SessionID:         uuid.New(),
		Registered_number: RegisteredNumber,
		Issued_at:         timeNow,
		Expired_at:        timeNow.Add(tokenDuration),
	}
}

func (pld *Payload) Valid() error {
	// if time.Now().After(pld.Expired_at) {
	// 	return jwt.ErrTokenExpired
	// }
	return nil
}
