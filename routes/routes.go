package routes

import (
	"net/http"
	"projects/webapi/api/users"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func AppMux() http.Handler {

	m := mux.NewRouter().StrictSlash(true)

	//Users routes

	u := m.PathPrefix("/users").Subrouter()

	u.
		Methods("GET").
		Path("/").
		Handler(alice.New(logMiddleware).ThenFunc(users.IndexHandler))

	u.
		Methods("POST").
		Path("/create").
		Handler(alice.New(logMiddleware).ThenFunc(users.CreateHandler))

	uid := m.PathPrefix("/users/{id}").Subrouter()

	uid.
		Methods("GET").
		Handler(alice.New(logMiddleware).ThenFunc(users.ShowHandler))
	uid.
		Methods("PUT").
		Handler(alice.New(logMiddleware).ThenFunc(users.UpdateHandler))

	return m

}
