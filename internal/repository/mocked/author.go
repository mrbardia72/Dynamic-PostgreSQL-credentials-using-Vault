package mocked

import (
	"context"
	"github.com/stretchr/testify/mock"
	"vault-psql/internal/service/author"
)

type MockAuthorRepository struct {
	mock.Mock
}

func NewMockAuthorRepository() *MockAuthorRepository {
	return &MockAuthorRepository{}
}

func (repo *MockAuthorRepository) List(ctx context.Context) ([]author.Entity, error) {
	args := repo.Called(ctx)

	return args.Get(0).([]author.Entity), args.Error(1) //nolint:wrapcheck
}
