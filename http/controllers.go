package http

import (
	"go-postr/html"
	"log"
	"net/http"
)

func homePageController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := "invalid method for home page: expected get"

		log.Println(response)
		http.Error(w, response, http.StatusNoContent)

		return
	}

	files := html.GetFiles("Partials.Base", "Home")

	tmpl, err := html.ParseFiles(files...)
	if err != nil {
		log.Println("could not parse files:", err)
		http.Error(w, err.Error(), http.StatusNoContent)

		return
	}

	data := struct{PartialBaseParams html.PartialsBaseParams}{
		PartialBaseParams: html.PartialsBaseParams{
			DisplayHeader: true,
		},
	}

	tmpl.Execute(w, r, "Partials.Base", data)
}
