package main

import (
	"net/http"
	"projects/webapi/routes"
)

func main() {
	http.ListenAndServe(":3000", routes.AppMux())
}
