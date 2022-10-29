package author

import "context"

type Service interface {
	ListAuthors(ctx context.Context) ([]Entity, error)
}
