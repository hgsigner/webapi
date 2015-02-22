package main

import (
	"log"
	"net/http"
	"projects/webapi/routes"
)

func main() {
	log.Println("Listening to port 3000")
	http.ListenAndServe(":3000", routes.AppMux())
}
