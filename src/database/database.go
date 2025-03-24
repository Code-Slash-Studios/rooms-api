package database

import (
	"database/sql"
	"log"
)

func dbConnect() *sql.DB {
	dsn := "code:dev</>@tcp(10.25.1.253:3306)/cis-rooms"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	log.Println("Connected to MariaDB successfully!")
	return db
}
