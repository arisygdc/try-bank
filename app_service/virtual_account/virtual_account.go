package virtualaccount

import (
	"context"

	"github.com/google/uuid"
)

func (svc Service) VirtualAccountID(ctx context.Context, va_identity int32) (uuid.UUID, error) {
	return svc.repos.Query().VirtualAccountID(ctx, va_identity)
}
