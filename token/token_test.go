package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJWToken(t *testing.T) {
	secret := "ad79ba5aa97d459af14a20cdbe57c60c5a50afa51a7df2a68073c22ce4783361"
	d, err := time.ParseDuration("3h")
	assert.Nil(t, err)

	payload := Newpayload(87423, d)
	tokenMaker, err := NewJWT(secret)
	assert.Nil(t, err)

	token, err := tokenMaker.Make(payload)
	assert.Nil(t, err)

	verifiedPayload, err := tokenMaker.Verify(token)
	assert.NoError(t, err, err)

	if assert.NotNil(t, verifiedPayload) {
		assert.True(t, verifiedPayload.Registered_number == payload.Registered_number, "session id not match")
	}
}
