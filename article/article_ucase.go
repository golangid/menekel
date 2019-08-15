package article

import (
	"context"
	"time"

	"github.com/golangid/menekel"
)

type articleUsecase struct {
	articleRepo    menekel.ArticleRepository
	contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of menekel.ArticleUsecase interface
func NewArticleUsecase(a menekel.ArticleRepository, timeout time.Duration) menekel.ArticleUsecase {
	if a == nil {
		panic("Article repository is nil")
	}
	if timeout == 0 {
		panic("Timeout is empty")
	}
	return &articleUsecase{
		articleRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *articleUsecase) Fetch(c context.Context, cursor string, num int64) (res []menekel.Article, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, nextCursor, err = a.articleRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	return
}

func (a *articleUsecase) GetByID(c context.Context, id int64) (res menekel.Article, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.articleRepo.GetByID(ctx, id)
	return
}

func (a *articleUsecase) Update(c context.Context, ar *menekel.Article) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.articleRepo.Update(ctx, ar)
}

func (a *articleUsecase) GetByTitle(c context.Context, title string) (res menekel.Article, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err = a.articleRepo.GetByTitle(ctx, title)
	return
}

func (a *articleUsecase) Store(c context.Context, m *menekel.Article) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, _ := a.GetByTitle(ctx, m.Title)
	if existedArticle != (menekel.Article{}) {
		return menekel.ErrConflict
	}

	err = a.articleRepo.Store(ctx, m)
	return
}

func (a *articleUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, err := a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedArticle == (menekel.Article{}) {
		return menekel.ErrNotFound
	}
	return a.articleRepo.Delete(ctx, id)
}
