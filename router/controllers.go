package router

import (
	"fmt"
	"go-postr/db"
	"go-postr/templates"
	"net/http"

	"github.com/charmbracelet/log"
)

func renderIndexController(w http.ResponseWriter, r *http.Request) {
	templ := templates.NewTemplateRenderer()

	templ.Render(w, "Index", nil)
}

func searchUsernameController(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	query := v.Get("query")

	if len(query) == 0 {
		fmt.Fprintf(w, "")  // insert an empty string in the results tag

		return
	}

	list, err := db.SearchByUsername(query)
	if err != nil {
		log.Error("Couldn't search for user " + query, "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	fmt.Fprintf(w, list)
}

func getRecentArticlesController(w http.ResponseWriter, r *http.Request) {
	conn := db.Connection()

	const limit = 5

	rows, _ := conn.Query(`SELECT u."username" as "author", a."content" as "article" FROM "article" a
		LEFT JOIN "user" u on u."id" = a."user_id" ORDER BY RANDOM() LIMIT $1`, limit)

	defer rows.Close()

	res := ""

	for rows.Next() {
		var username, article string
		
		_ = rows.Scan(&username, &article)

		res += fmt.Sprintf(`<li>
			<em>"%s"</em> &mdash; <a href="#">%s</a>
		</li>`, article, username)
	}

	fmt.Fprintf(w, res)
}
