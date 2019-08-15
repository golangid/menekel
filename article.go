package menekel

import (
	"context"
	"time"
)

// Article represent the Article contract
type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ArticleRepository represent the repository contract
type ArticleRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Article, csr string, err error)
	GetByID(ctx context.Context, id int64) (res Article, err error)
	GetByTitle(ctx context.Context, title string) (res Article, err error)
	Update(ctx context.Context, article *Article) (err error)
	Store(ctx context.Context, a *Article) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
