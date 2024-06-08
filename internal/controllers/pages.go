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
	users_basic_info, _ := pc.db_service.RecentlyCreatedUsers(6)

	pages.HomePage(users_basic_info).Render(r.Context(), w)
}
