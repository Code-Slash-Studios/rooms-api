package main

import (
	"database/sql"
	_ "fmt"
	"log"
	"net/http"
	_ "rooms-api/src/database"
	"rooms-api/src/router"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Init DataBase Connection
	//db = database.dbConnect()
	//defer db.Close()

	router := router.Router()
	log.Println("Server is listening on port 6000")
	http.ListenAndServe(":6000", router)
}
