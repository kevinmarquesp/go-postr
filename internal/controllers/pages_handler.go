package controllers

import (
	"net/http"

	"github.com/kevinmarquesp/go-postr/views/pages"
)

type PagesHandler struct {
}

func NewPagesHanlder() PagesHandler {
	return PagesHandler{}
}

func (ph PagesHandler) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(r.Context(), w)
}
