package http

import (
	"go-postr/html"
	"log"
	"net/http"
)

func homePageController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed: expected get", http.StatusMethodNotAllowed)
		return
	}

	files := html.GetFiles("Partials.Base", "Home")

	tmpl, err := html.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct{PartialBaseParams html.PartialsBaseParams}{
		PartialBaseParams: html.PartialsBaseParams{
			DisplayHeader: true,
		},
	}

	tmpl.Execute(w, r, "Partials.Base", data)
}

func signupPageController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed: expected get", http.StatusMethodNotAllowed)
		return
	}

	files := html.GetFiles("Partials.Base", "Signup")

	tmpl, err := html.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := struct{PartialBaseParams html.PartialsBaseParams}{
		PartialBaseParams: html.PartialsBaseParams{
			DisplayHeader: true,
		},
	}

	tmpl.Execute(w, r, "Partials.Base", data)
}

func usernameValidationController(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	username := r.Form.Get("username")
	log.Println("username:", username)
}

func getFieldValidationStatusComponent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	message := r.Form.Get("message")
	bsStatus := r.Form.Get("bs_status")

	log.Println("status msg:", message)
	log.Println("status bsStatus:", bsStatus)

	files := html.GetFiles("Components.FieldValidationStatus")
	tmpl, _ := html.ParseFiles(files...)

	data := struct{Params html.ComponentsFieldvalidationstatusParams}{
		Params: html.ComponentsFieldvalidationstatusParams{
			BootstrapStatus: "danger",
			Message: "Hello world, This is a cool error!",
		},
	}

	tmpl.Execute(w, r, "Components.FieldValidationStatus", data)
}