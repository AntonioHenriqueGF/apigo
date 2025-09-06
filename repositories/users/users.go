package users

import (
	"context"

	"github.com/AntonioHenriqueGF/apigo/schemas"
)

type UsersRepositorieInterface interface {
	Create(ctx context.Context, user *schemas.User) error
	GetByID(ctx context.Context, id string) (*schemas.User, error)
	Login(ctx context.Context, email string, password string) error
	Update(ctx context.Context, user *schemas.User) error
	Delete(ctx context.Context, id string) error
}
