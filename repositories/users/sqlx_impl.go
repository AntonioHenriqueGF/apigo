package users

import (
	"context"
	"fmt"

	"github.com/AntonioHenriqueGF/apigo/schemas"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type userDbSqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSqlx(w *sqlx.DB, r *sqlx.DB) UsersRepositorieInterface {
	return &userDbSqlx{
		writer: w,
		reader: r,
	}
}

func (u *userDbSqlx) Create(ctx context.Context, user *schemas.User) error {
	existingUser, err := u.GetByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	password := user.Password

	// Generate a bcrypt hash of the password.
	// bcrypt automatically generates a unique salt for each hash.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	_, err = u.writer.NamedExecContext(ctx, `INSERT INTO users (usu_name, usu_email, usu_password_hash) VALUES (:usu_name, :usu_email, :usu_password)`, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userDbSqlx) GetByEmail(ctx context.Context, email string) (*schemas.UserBD, error) {
	var user schemas.UserBD
	err := u.reader.GetContext(ctx, &user, `SELECT * FROM users WHERE usu_email = ?`, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userDbSqlx) Login(ctx context.Context, email string, password string) (*schemas.UserBD, error) {
	var user schemas.UserBD
	err := u.reader.GetContext(ctx, &user, `SELECT * FROM users WHERE usu_email = ?`, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userDbSqlx) GetByID(ctx context.Context, id string) (*schemas.User, error) {
	var user schemas.User
	err := u.reader.GetContext(ctx, &user, `SELECT * FROM users WHERE usu_id = ?`, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userDbSqlx) Update(ctx context.Context, user *schemas.User) error {
	// Implementation goes here
	return nil
}

func (u *userDbSqlx) Delete(ctx context.Context, id string) error {
	// Implementation goes here
	return nil
}
