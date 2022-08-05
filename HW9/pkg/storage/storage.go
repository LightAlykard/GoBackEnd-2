package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/lightlykard/gobackend-2/hw9/pkg/models"
)

type Storage interface {
	Create(ctx context.Context, list models.List) error
	Read(ctx context.Context, id uuid.UUID) (list *models.List, err error)
	Update(ctx context.Context, id uuid.UUID, items []*models.Item) error
	Delete(ctx context.Context, id uuid.UUID) error
}
