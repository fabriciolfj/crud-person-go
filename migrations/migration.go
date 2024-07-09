package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Failed to create driver: %v", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not get current directory: %v", err)
	}

	// Construct the path to migration scripts
	migrationPath := fmt.Sprintf("file://%s/migrations/scripts", dir)

	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"db", driver)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	// Only call Steps if Migrate instance was successfully created
	err = m.Up()
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
