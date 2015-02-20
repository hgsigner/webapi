package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/Schema"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	uRenderer *render.Render
)

func init() {
	uRenderer = render.New()
}

type User struct {
	ID        bson.ObjectId `bson:"_id" json:"_id" schema:"_id"`
	FirstName string        `bson:"fisrt_name" json:"first_name" schema:"first_name"`
	LastName  string        `bson:"last_name" json:"last_name" schema:"last_name"`
	Email     string        `bson:"email" json:"email" schema:"email"`
}

type UserResults struct {
	Users []User `json:"users"`
}

func uCollection(session *mgo.Session) *mgo.Collection {
	return session.DB(AuthDatabase).C("users")
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

	newUser := User{}
	err = collection.Find(bson.M{"_id": u.ID}).One(&newUser)
	if err != nil {
		log.Fatal(err)
	}

	uRenderer.JSON(w, http.StatusCreated, UserResults{[]User{newUser}})

}

func userMux() http.Handler {

	m := mux.NewRouter()
	m.HandleFunc("/users", wellcomeHandler).Methods("GET")
	m.HandleFunc("/users/create", createUserHandler).Methods("POST")
	return m

}
