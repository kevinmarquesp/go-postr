package controllers

import (
	"net/http"

	"github.com/kevinmarquesp/go-postr/internal/models"
	"github.com/kevinmarquesp/go-postr/views/pages"
)

type PagesController struct {
	db_service models.DatabaseService
}

func NewPagesController(db_service models.DatabaseService) PagesController {
	return PagesController{
		db_service: db_service,
	}
}

func (pc PagesController) HomePage(w http.ResponseWriter, r *http.Request) {
	pages.HomePage().Render(r.Context(), w)
}
