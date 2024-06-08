package models

// This interface is the type that the router, controllers, and other parts of
// the code, will use to communicate to the database. It only allows to use
// the methods listed below, delegating the responsabilite to handle each
// service (such as Postgres, SQLite3 or something else) to the `main()` func.
type DatabaseService interface {
	// Connect to the database server. The credentials will be fetched from
	// the systems's environment variables in the implementation of this
	// function.
	Connect() error

	// TODO: Add a documentation for this function
	InsertUser(username, description, password string) error

	// TODO: Add a documentation for this function
	RecentlyCreatedUsers(size int) ([]UserBasicInfo, error)
}
