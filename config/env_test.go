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
	assert.Equal(t, "debug", env.Environment)
}
