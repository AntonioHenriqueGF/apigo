package repositories

import (
	"github.com/AntonioHenriqueGF/apigo/repositories/posts"
	"github.com/jmoiron/sqlx"
)

type Options struct {
	WriterSqlx *sqlx.DB
	ReaderSqlx *sqlx.DB
}

type Container struct {
	Post posts.PostRepositorieInterface
}

func New(options Options) *Container {
	return &Container{
		Post: posts.NewSqlx(options.ReaderSqlx, options.WriterSqlx),
	}
}
