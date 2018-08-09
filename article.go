package menekel

import (
	"context"
	"time"
)

type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Author    Author    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
type ArticleRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*Article, error)
	GetByID(ctx context.Context, id int64) (*Article, error)
	GetByTitle(ctx context.Context, title string) (*Article, error)
	Update(ctx context.Context, article *Article) (*Article, error)
	Store(ctx context.Context, a *Article) (int64, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
