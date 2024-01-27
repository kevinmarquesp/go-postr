package controllers

import (
	"database/sql"
	"go-postr/models"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreateNewUserController(w http.ResponseWriter, r *http.Request) {
	if err := parseValidationFormFields(w, r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[ERROR]", err)

		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	var userBio *sql.NullString

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(username + password), bcrypt.MinCost)
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}

	_, err = models.Db.Exec(`INSERT INTO public.user (username, password, bio, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $4)`, username, hashedPassword, userBio, time.Now())
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
