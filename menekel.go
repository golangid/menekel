package menekel

import "context"

type ArticleUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*Article, string, error)
	GetByID(ctx context.Context, id int64) (*Article, error)
	Update(ctx context.Context, ar *Article) (*Article, error)
	GetByTitle(ctx context.Context, title string) (*Article, error)
	Store(context.Context, *Article) (*Article, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
