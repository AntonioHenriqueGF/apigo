package posts

import (
	"context"

	"github.com/AntonioHenriqueGF/apigo/schemas"
)

// PostRepositorieInterface is the interface that wraps the basic methods for the post repository.
type PostRepositorieInterface interface {
	GetAll(ctx context.Context) ([]schemas.Post, error)
}
