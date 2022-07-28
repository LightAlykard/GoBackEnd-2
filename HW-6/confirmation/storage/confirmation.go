package storage

import (
	"context"

	"github.com/LightAlykard/GoBackEnd-2/HW6/models"
	//"models"
)

type ConfirmationStorage interface {
	Create(ctx context.Context, c *models.Confirmation) error
	GetConfirmation(ctx context.Context, name string) (*models.Confirmation, error)
	Delete(ctx context.Context, name string) error
}
