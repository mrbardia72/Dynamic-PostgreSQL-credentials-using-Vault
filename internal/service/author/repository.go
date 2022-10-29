package author

import "context"

type Repository interface {
	List(ctx context.Context) ([]Entity, error)
}
