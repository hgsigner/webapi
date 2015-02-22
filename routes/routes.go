package routes

import (
	"net/http"
	"projects/webapi/api/users"

	"github.com/gorilla/mux"
)

func AppMux() http.Handler {

	m := mux.NewRouter().StrictSlash(true)

	//Users routes

	u := m.PathPrefix("/users").Subrouter()
	u.Methods("GET").Path("/").HandlerFunc(users.IndexHandler)
	u.Methods("GET").Path("/{id}").HandlerFunc(users.ShowHandler)
	u.Methods("POST").Path("/create").HandlerFunc(users.CreateHandler)

	return m

}
