//go:build author_logic_test

package author_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"vault-psql/internal/repository/mocked"
	"vault-psql/internal/service/author"
)

func authorService() (context.Context, *mocked.MockAuthorRepository, author.Service) {
	ctx := context.Background()
	authorRepo := mocked.NewMockAuthorRepository()
	authorSvc := author.NewService(authorRepo)
	return ctx, authorRepo, authorSvc
}

func TestService_ListAuthors(t *testing.T) {
	ctx, authorRepo, authorSvc := authorService()
	t.Run("test list author", func(t *testing.T) {
		var res []author.Entity

		authorRepo.On("List", ctx).Return(res, nil)

		_, err := authorSvc.ListAuthors(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, http.StatusOK)
	})
}
