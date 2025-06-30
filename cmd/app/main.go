package app

import (
	"database/sql"
	"log"
)

func main() {
	connStr := "postgres://user:password@localhost:5432/dbname?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()
}
