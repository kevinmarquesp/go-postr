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

	files := html.GetFiles("Components.FieldValidationStatus")
	tmpl, _ := html.ParseFiles(files...)

	//todo: check if the user is valid and if its already exists on the database
	data := struct{Params html.ComponentsFieldvalidationstatusParams}{
		Params: html.ComponentsFieldvalidationstatusParams{
			BootstrapStatus: "success",
			Message: "You're username looks unique, good job!",
			// BootstrapStatus: "danger",
			// Message: "Username already taken",
		},
	}

	// w.WriteHeader(http.StatusConflict)
	tmpl.Execute(w, r, "Components.FieldValidationStatus", data)
}

//todo: add a "controller" at the end of the name of this function
func getFieldValidationStatusComponent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	message := r.Form.Get("message")
	bootstrapStatus := r.Form.Get("bootstrap-status")

	log.Println("status msg:", message)
	log.Println("status bsStatus:", bootstrapStatus)

	files := html.GetFiles("Components.FieldValidationStatus")
	tmpl, _ := html.ParseFiles(files...)

	data := struct{Params html.ComponentsFieldvalidationstatusParams}{
		Params: html.ComponentsFieldvalidationstatusParams{
			BootstrapStatus: bootstrapStatus,
			Message: message,
		},
	}

	tmpl.Execute(w, r, "Components.FieldValidationStatus", data)
}

func createNewUserController(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	log.Println("username:", username)
	log.Println("password:", password)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
