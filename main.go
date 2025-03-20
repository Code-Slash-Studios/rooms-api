package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"rooms-api/src/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "code:dev</>@tcp(0.0.0.0:3306)/cis-rooms"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	log.Println("Connected to MariaDB successfully!")

	router := router.Router()
	fmt.Println("Server is listening on port 6000")
	http.ListenAndServe(":6000", router)
}
