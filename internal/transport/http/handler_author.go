package dockertest_test

import (
	"encoding/json"
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

func (h *authorHandler) HandlerListAuthors(w http.ResponseWriter, r *http.Request) {
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

func (h *authorHandler) HandlerCreateAuthor(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Name string `json:"name" yaml:"name"`
	}

	type Response struct {
		ID int `json:"id" yaml:"id"`
	}

	var req Request

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		respond.Done(w, r, problem.BadRequest("invalid request payload"))

		return
	}

	idAuthors, err := h.authorSvc.CreateAuthor(r.Context(), author.CreateAuthorRequest{
		Name: req.Name})
	if err != nil {
		respond.Done(w, r, problem.InternalServerError(err))
	}

	respond.Done(w, r, Response{ID: idAuthors})
}
