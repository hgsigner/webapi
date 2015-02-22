package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projects/webapi/db"

	"github.com/gorilla/Schema"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `bson:"_id" json:"_id" schema:"_id"`
	FirstName string        `bson:"fisrt_name" json:"first_name" schema:"first_name"`
	LastName  string        `bson:"last_name" json:"last_name" schema:"last_name"`
	Email     string        `bson:"email" json:"email" schema:"email"`
}

type UserResults struct {
	Users []User `json:"users"`
}

type JsonHttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Welcome!")
}

func ShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-Type", "application/json")

	session := db.InitDb().Copy()
	defer session.Close()
	collection := db.GetCollection(session, "users")

	user := User{}
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(JsonHttpError{404, "Not Found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UserResults{[]User{user}})
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	session := db.InitDb().Copy()
	defer session.Close()

	collection := db.GetCollection(session, "users")

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

	newUser := User{}
	err = collection.Find(bson.M{"_id": u.ID}).One(&newUser)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(UserResults{[]User{newUser}})

}
