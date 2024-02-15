package posts

import (
	"context"

	"github.com/AntonioHenriqueGF/apigo/schemas"
	"github.com/jmoiron/sqlx"
)

type postDbSqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSqlx(w *sqlx.DB, r *sqlx.DB) PostRepositorieInterface {
	return &postDbSqlx{
		writer: w,
		reader: r,
	}
}

func (p *postDbSqlx) GetAll(ctx context.Context) ([]schemas.Post, error) {
	var posts []schemas.Post

	err := p.reader.SelectContext(ctx, &posts, `SELECT * FROM posts`)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *postDbSqlx) Create(ctx context.Context, post *schemas.Post) error {
	_, err := p.writer.NamedExecContext(ctx, `INSERT INTO posts (title, content) VALUES (:title, :content)`, post)

	if err != nil {
		return err
	}

	return nil
}
