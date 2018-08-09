package author

import (
	"context"

	models "github.com/golangid/menekel"
)

type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (*models.Author, error)
}
