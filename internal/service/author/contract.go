package author

import "context"

type Service interface {
	ListAuthors(ctx context.Context) ([]Entity, error)
	CreateAuthor(ctx context.Context, req CreateAuthorRequest) (id int, err error)
}

type Repository interface {
	List(ctx context.Context) ([]Entity, error)
	Insert(ctx context.Context, entity Entity) (id int, err error)
}

type CreateAuthorRequest struct {
	Name string
}
