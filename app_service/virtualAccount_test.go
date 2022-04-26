package appservice

import (
	"context"
	"testing"
	virtualaccount "try-bank/app_service/virtual_account"
	"try-bank/config"
	"try-bank/database"

	"github.com/google/uuid"
)

func TestVirtualaccount(t *testing.T) {
	ctx := context.Background()
	ctx, ctxCancel := context.WithCancel(ctx)

	defer ctxCancel()
	repo, err := database.NewRepository(config.Environment{
		DBDriver: database.PostgreDriver,
		DBSource: "postgresql://postgres:secret@localhost:5432/bank?sslmode=disable",
	}, ctx)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	svc := virtualaccount.New(repo)
	_, err = svc.Register(ctx, uuid.New(), "http://kon.trol")
	if err == nil {
		t.Error(err)
		t.Fail()
	}
}
