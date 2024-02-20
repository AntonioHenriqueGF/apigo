package posts

import (
	"context"

	"github.com/AntonioHenriqueGF/apigo/schemas"
)

// PostRepositorieInterface is the interface that wraps the basic methods for the post repository.
type PostRepositorieInterface interface {
	GetAll(ctx context.Context) ([]schemas.Post, error)
	Create(ctx context.Context, post *schemas.Post) error
	ExistsByID(ctx context.Context, id string) (bool, error)
	GetByID(ctx context.Context, id string) (*schemas.Post, error)
	DeleteByID(ctx context.Context, id string) error
}
