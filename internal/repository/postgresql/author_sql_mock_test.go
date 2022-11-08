//go:build sql_mock_author

package postgresql

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

// a successful case
func TestShouldUpdateStats(t *testing.T) {
	db, mock := NewMock()
	repo := &authorRepo{db}
	defer db.Close()

	rows := mock.NewRows([]string{"author_id", "first_name", "last_name"}).
		AddRow(1, "post1", "hello").
		AddRow(2, "post2", "world")
	mock.ExpectQuery("^SELECT author_id,first_name,last_name FROM authors$").WillReturnRows(rows)

	user, err := repo.List(context.Background())
	assert.NotNil(t, user)
	assert.NoError(t, err)

	// ExpectationsWereMet checks whether all queued expectations were met in order.
	if errExpectation := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", errExpectation)
	}
}
