package db

import "go-postr/templates"

//  TODO: Make these functions return palain golang object types
//  TODO: Create a separate file that will build the HTML response strings

func SearchByUsername(query string) (string, error) {
	rows, err := conn.Query(`SELECT username FROM "user" WHERE username LIKE $1`, "%"+query+"%")
	if err != nil {
		return "", err
	}

	defer rows.Close()

	templ := templates.NewTemplateRenderer()
	res := ""
	i := 0

	for rows.Next() {
		var username string

		err = rows.Scan(&username)
		if err != nil {
			return "", err
		}

		props := &templates.UsernameSearchItemResultComponentProps{
			IsSelected: i == 0,
			Username:   username,
		}

		resElm, err := templ.RenderString("UsernameSearchItemResult.Component", props)
		if err != nil {
			return "", err
		}

		res += resElm
		i++
	}

	return res, nil
}

// No, it will not fetch the most recent articles, but it will look like it does
// that. And that's enough...
func GetRecentArticles() (string, error) {
	const limit = 5

	rows, err := conn.Query(`SELECT u."username" as "author", a."content" as "article" FROM "article" a
		LEFT JOIN "user" u on u."id" = a."user_id" ORDER BY RANDOM() LIMIT $1`, limit)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	templ := templates.NewTemplateRenderer()
	res := ""

	for rows.Next() {
		var author, article string

		err = rows.Scan(&author, &article)
		if err != nil {
			return "", err
		}

		props := &templates.ArticleCardComponentProps{
			Article: article,
			Author:  author,
		}

		resElm, err := templ.RenderString("ArticleCard.Component", props)
		if err != nil {
			return "", err
		}

		res += resElm
	}

	return res, nil
}

func WasUsernameAlreadyTaken(username string) (bool, error) {
	rows, err := conn.Query(`SELECT username FROM "user" WHERE username LIKE $1`, username)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	for rows.Next() {
		var dbusername string

		err = rows.Scan(&dbusername)
		if err != nil {
			return false, err
		}

		if dbusername == username {
			return true, nil
		}
	}

	return false, nil
}
