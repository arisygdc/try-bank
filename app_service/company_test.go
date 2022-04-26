package appservice

import (
	"context"
	"testing"
	"try-bank/app_service/company"
	"try-bank/database"

	"github.com/stretchr/testify/assert"
)

func TestCreateCompany(t *testing.T) {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	env := getDBConf()

	repos, err := database.NewRepository(env, ctx)
	assert.Nil(t, err)

	register := company.RegisterCompanyParam{
		Name:    "company companyan",
		Email:   "mail@company.com",
		Phone:   "081638293643",
		Pin:     "083682",
		Deposit: 30000000,
	}

	svc := company.New(repos)

	registered, err := svc.CreateCompanyAccount(ctx, register)
	assert.Nil(t, err)
	t.Log(registered)

	_, err = svc.CreateCompanyAccount(ctx, register)
	assert.Error(t, err)
}
