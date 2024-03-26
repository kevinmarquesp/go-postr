package db

import "fmt"

func SearchByUsername(query string) (string, error) {
	res := ""

	rows, err := conn.Query(`SELECT username FROM "user" WHERE username LIKE $1`, "%" + query + "%")
	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		var username string

		err = rows.Scan(&username)
		if err != nil {
			return "", err
		}

		res += fmt.Sprintf(`<li>
			<a href="#">%s</a>
		</li>`, username)
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

	res := ""

	for rows.Next() {
		var username, article string
		
		err = rows.Scan(&username, &article)
		if err != nil {
			return "", err
		}

		res += fmt.Sprintf(`<li>
			<em>"%s"</em> &mdash; <a href="#">%s</a>
		</li>`, article, username)
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
