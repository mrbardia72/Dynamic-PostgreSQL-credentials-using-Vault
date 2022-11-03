package postgresql

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"vault-psql/internal/service/author"
)

const (
	queryList = `SELECT author_id,first_name,last_name FROM authors`
)

type authorRepo struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) author.Repository {
	return &authorRepo{db: db}
}

//List query from authors.
func (repo *authorRepo) List(ctx context.Context) ([]author.Entity, error) {
	//tx, err := repo.db.Begin()
	//if err != nil {
	//	return nil, nil
	//}
	//
	//defer func() {
	//	switch err {
	//	case nil:
	//		err = tx.Commit()
	//		log.Fatalf("update drivers: unable to commit: %v", err)
	//	default:
	//		if rollbackErr := tx.Rollback(); rollbackErr != nil {
	//			log.Fatalf("update drivers: unable to rollback: %v", rollbackErr)
	//		}
	//	}
	//}()
	rows, err := repo.db.QueryContext(ctx, queryList)
	if err != nil {
		return nil, errors.Wrap(err, "error on query list author")
	}

	defer func() { _ = rows.Close() }()

	actuators := make([]author.Entity, 0)

	for rows.Next() {
		var p author.Entity

		if errScan := rows.Scan(&p.ID, &p.FirstName, &p.LastName); errScan != nil {
			return nil, errors.Wrap(errScan, "error on query List scan row")
		}

		actuators = append(actuators, p)
	}

	return actuators, nil
}
