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

// NewSqlx creates a new instance of the post repository using sqlx.
func NewSqlx(w *sqlx.DB, r *sqlx.DB) PostRepositorieInterface {
	return &postDbSqlx{
		writer: w,
		reader: r,
	}
}

// GetAll returns all posts from the database.
func (p *postDbSqlx) GetAll(ctx context.Context) ([]schemas.Post, error) {
	var posts []schemas.Post

	err := p.reader.SelectContext(ctx, &posts, `SELECT * FROM posts`)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

// Create creates a new post in the database.
func (p *postDbSqlx) Create(ctx context.Context, post *schemas.Post) error {
	result, err := p.writer.NamedExecContext(ctx, `INSERT INTO posts (pst_title, pst_content) VALUES (:pst_title, :pst_content)`, post)

	if err != nil {
		return err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	post.ID = int(lastId)

	return nil
}

// GetByID returns a post from the database by its ID.
func (p *postDbSqlx) GetByID(ctx context.Context, id string) (*schemas.Post, error) {
	post := &schemas.Post{}

	err := p.reader.GetContext(ctx, post, `SELECT * FROM posts WHERE pst_id = ?`, id)

	if err != nil {
		return nil, err
	}

	return post, nil
}
