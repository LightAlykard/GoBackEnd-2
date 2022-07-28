package storage

import (
	"context"

	"github.com/LightAlykard/GoBackEnd-2/HW-6/models"
)

type UserStorage interface {
	Create(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, name string) (*models.User, error)
}
