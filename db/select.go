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
