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
