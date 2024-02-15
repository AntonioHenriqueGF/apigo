package repositories

import (
	"github.com/AntonioHenriqueGF/apigo/repositories/posts"
	"github.com/jmoiron/sqlx"
)

type Options struct {
	WriterSqlx *sqlx.DB
	ReaderSqlx *sqlx.DB
}

// Container holds all repositories.
type Container struct {
	Post posts.PostRepositorieInterface
}

// New creates a new instance of the repositories container.
func New(options Options) *Container {
	return &Container{
		Post: posts.NewSqlx(options.ReaderSqlx, options.WriterSqlx),
	}
}
