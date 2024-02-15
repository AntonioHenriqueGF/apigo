package posts

import (
	"context"

	"github.com/AntonioHenriqueGF/apigo/schemas"
)

type PostRepositorieInterface interface {
	GetAll(ctx context.Context) ([]schemas.Post, error)
}
