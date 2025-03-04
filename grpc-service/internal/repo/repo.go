package repo

import (
	"context"

	"github.com/artromone/lccs/grpc-service/internal/models"
)

type BookRepository interface {
	Create(ctx context.Context, book *models.Book) error
	Get(ctx context.Context, id string) (*models.Book, error)
}
