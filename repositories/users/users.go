package users

import (
	"context"

	"github.com/AntonioHenriqueGF/apigo/schemas"
)

type UsersRepositorieInterface interface {
	CreateUser(ctx context.Context, user *schemas.User) error
	GetUserByID(ctx context.Context, id string) (*schemas.User, error)
	UpdateUser(ctx context.Context, user *schemas.User) error
	DeleteUser(ctx context.Context, id string) error
}
