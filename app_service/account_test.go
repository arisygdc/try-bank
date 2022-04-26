package appservice

import (
	"context"
	"testing"
	"time"
	"try-bank/app_service/account"
	"try-bank/database"

	"github.com/stretchr/testify/assert"
)

func TestUserAccount(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	env := getDBConf()
	defer cancel()

	repos, err := database.NewRepository(env, ctx)
	assert.Nil(t, err)

	svc := account.New(repos)
	accountType, err := svc.GetAccountType(ctx, account.LevelSilver)
	assert.Nil(t, err, "cannot get account type")

	register := account.CreateUserParam{
		Firstname:   "arisy",
		Lastname:    "musyafa'",
		Email:       "arisy@gdc.com",
		Phone:       "081217827013",
		Pin:         "025361",
		Birth:       time.Now(),
		AccountType: accountType.ID,
	}
	_, err = svc.CreateCustomerAccount(ctx, register)
	assert.Nil(t, err, "error create account")
	_, err = svc.CreateCustomerAccount(ctx, register)
	assert.Error(t, err)
}
