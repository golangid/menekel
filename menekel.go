package menekel

import "context"

//ArticleUsecase represent the usecase of the article
type ArticleUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Article, csr string, err error)
	GetByID(ctx context.Context, id int64) (res Article, err error)
	Update(ctx context.Context, ar *Article) (err error)
	GetByTitle(ctx context.Context, title string) (res Article, err error)
	Store(context.Context, *Article) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
