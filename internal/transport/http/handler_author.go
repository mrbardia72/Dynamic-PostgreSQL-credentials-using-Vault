package dockertest_test

import (
	"net/http"
	"vault-psql/internal/service/author"
	"vault-psql/pkg/problem"
	"vault-psql/pkg/respond"
)

type authorHandler struct {
	authorSvc author.Service
}

func newAuthorHandler(authorApp author.Service) *authorHandler {
	return &authorHandler{
		authorSvc: authorApp,
	}
}

func (h *authorHandler) ListAuthors(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Results []author.Entity `json:"results"`
	}

	res, err := h.authorSvc.ListAuthors(r.Context())
	if err != nil {
		respond.Done(w, r, problem.InternalServerError(err))

		return
	}

	respond.Done(w, r, Response{Results: res})

	return
}
