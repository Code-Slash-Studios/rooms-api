package main

import (
	"fmt"
	"net/http"
	"rooms-api/src/routes"
)

func main() {
	router := routes.Router()
	fmt.Println("Server is listening on port 6000")
	http.ListenAndServe(":6000", router)
}
