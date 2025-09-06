package store

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib" // "_" underscore will not let formatter to remove the import
)

/*
1. db server runnning
2. we need to connect to db server
3. we need some middlerware (here driver) to do that
4. some config to open and close db connections
*/

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")

	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	fmt.Println("Connected to db...")
	return db, nil
}
