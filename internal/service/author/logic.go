package author

import (
	"context"
	"github.com/pkg/errors"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (svc *service) ListAuthors(ctx context.Context) ([]Entity, error) {
	res, err := svc.repo.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error on list authors")
	}

	return res, nil
}

func (svc *service) CreateAuthor(ctx context.Context, req CreateAuthorRequest) (id int, err error) {
	m := Entity{
		Name: req.Name,
	}

	idAuthor, err := svc.repo.Insert(ctx, m)
	if err != nil {
		return 0, errors.Wrap(err, "error on create authors")
	}

	return idAuthor, nil
}
