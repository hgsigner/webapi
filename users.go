package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/Schema"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func uCollection(session *mgo.Session) *mgo.Collection {
	return session.DB(AuthDatabase).C("users")
}

type User struct {
	ID        bson.ObjectId `bson:"_id" json:"_id" schema:"_id"`
	FirstName string        `bson:"fisrt_name" json:"first_name" schema:"first_name"`
	LastName  string        `bson:"last_name" json:"last_name" schema:"last_name"`
	Email     string        `bson:"email" json:"email" schema:"email"`
}

func wellcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Welcome!")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	session := InitDb().Copy()
	defer session.Close()

	collection := uCollection(session)

	u := &User{}
	u.ID = bson.NewObjectId()
	decoder := schema.NewDecoder()
	err = decoder.Decode(u, r.Form)
	if err != nil {
		log.Fatal(err)
	}

	err = collection.Insert(u)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// data, _ := json.Marshal(u)
	// w.Write(data)

}

func userMux() http.Handler {

	m := mux.NewRouter()
	m.HandleFunc("/users", wellcomeHandler).Methods("GET")
	m.HandleFunc("/users/create", createUserHandler).Methods("POST")
	return m

}
