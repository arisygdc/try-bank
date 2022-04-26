package appservice

import (
	"context"
	"testing"
	"time"
	"try-bank/app_service/account"
	"try-bank/config"
	"try-bank/database"
)

func TestUserAccount(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	env, err := config.NewEnv("../")
	defer cancel()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	repos, err := database.NewRepository(env, ctx)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	svc := account.New(repos)
	lvlUser, err := svc.GetLevel(ctx, account.LevelClientUser)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	svc.CreateUserAccount(ctx, account.CreateUserParam{
		Firstname: "arisy",
		Lastname:  "musyafa'",
		Email:     "arisy@gdc.com",
		Phone:     "081217827013",
		Pin:       "025361",
		Birth:     time.Now(),
		Level:     lvlUser.Name,
	})
}
