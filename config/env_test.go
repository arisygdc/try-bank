package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	mp := make(map[string]string)
	mp["0.0.0.0:8080"] = "server address"
	mp["postgres"] = "database driver"
	mp["postgresql://root:secret@127.0.0.1:5432/bank?sslmode=disable"] = "database source"
	mp["debug"] = "environment"

	assert := assert.New(t)
	env, err := NewEnv("../")
	assert.NoError(err)

	h, ok := mp[env.DBDriver]
	assert.True(ok, h)

	h, ok = mp[env.DBSource]
	assert.True(ok, h)

	h, ok = mp[env.Environment]
	assert.True(ok, h)

	h, ok = mp[env.ServerAddress]
	assert.True(ok, h)
}
