package user

import (
	"context"

	"github.com/google/uuid"
)

type UserInteractor interface {
	AddCommunity(ctx context.Context, userID uuid.UUID, communityID uuid.UUID) (*User, error)
	DeleteCommunity(ctx context.Context, userID uuid.UUID, communityID uuid.UUID) (*User, error)
}
