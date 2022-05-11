package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	env, err := NewEnv("../.", "example.config")
	assert.NoError(t, err)

	assert.Equal(t, "0.0.0.0:8080", env.ServerAddress)
	assert.Equal(t, "postgres", env.DBDriver)
	assert.Equal(t, "postgresql://postgres:secret@localhost:5432/bank?sslmode=disable", env.DBSource)
	assert.Equal(t, "9cbd68bd99b4116e65a03e198c0b5241e60f5af41b4c5afbb7d6e7a30c982da6", env.TokenSymetricKey)
	assert.Equal(t, "debug", env.Environment)
}
