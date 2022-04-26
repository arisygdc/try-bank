package account

import (
	"context"

	"github.com/google/uuid"
)

type Level struct {
	ID   uuid.UUID
	Name string
}

func (svc Service) GetLevel(ctx context.Context, name_level string) (Level, error) {
	getLevel, err := svc.repos.Query().GetLevel(ctx, name_level)
	return Level(getLevel), err
}
