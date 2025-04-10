package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	_ "fmt"
	"log"
	"net/http"
	"rooms-api/src/router"
	"rooms-api/src/routes"
)

func main() {
	// Init DataBase Connection
	dsn := "code:dev</>@tcp(10.25.1.253:3306)/cis-rooms?parseTime=true&loc=Local"
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

	routes.DB = db

	router := router.Router()
	log.Println("Server is listening on port 6000")
	http.ListenAndServe(":6000", router)
}
