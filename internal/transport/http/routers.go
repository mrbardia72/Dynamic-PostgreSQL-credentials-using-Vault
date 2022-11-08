package dockertest_test

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"vault-psql/internal/service/author"
)

type ConfigService struct {
	Port      string
	AuthorSvc author.Service
}

func HandlerRequest(cfg ConfigService) error {
	authorHndlr := newAuthorHandler(cfg.AuthorSvc)

	r := mux.NewRouter()

	r.HandleFunc("/authors", authorHndlr.ListAuthors).Methods(http.MethodGet)

	err := http.ListenAndServe(cfg.Port, r) //host+port
	if err != nil {
		return errors.Wrap(err, "error on listen and serve http")
	}

	return nil
}
