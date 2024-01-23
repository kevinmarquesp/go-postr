package models

import (
	"log"
	"os"
)

var migrationFileLocation string = "models/migration.sql"

func ExecuteMigration() {
	file, err := os.ReadFile(migrationFileLocation)
	if err != nil {
		log.Println("Could not open the migration sql file")
		log.Panic(err)
	}

	_, err = Db.Exec(string(file))
	if err != nil {
		log.Println("Could not run the migraion sql query")
		log.Panic(err)
	}
}
