// database/db.go
package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("postgres", "host=localhost user=postgres password=yourpassword dbname=ecommerce sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}
