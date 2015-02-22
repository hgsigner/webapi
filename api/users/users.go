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
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id" schema:"_id"`
	FirstName string        `bson:"fisrt_name,omitempty" json:"first_name" schema:"first_name"`
	LastName  string        `bson:"last_name,omitempty" json:"last_name" schema:"last_name"`
	Email     string        `bson:"email,omitempty" json:"email" schema:"email"`
}

type UserResults struct {
	Users []User `json:"users"`
}

type JsonHttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Methods

func CreateUser(user User) (User, error) {

	session := db.InitDb().Copy()
	defer session.Close()
	collection := db.GetCollection(session, "users")

	u := User{}
	err := collection.Insert(user)
	if err != nil {
		return u, err
	}

	f_user, f_err := FindUser(user.ID.Hex())
	if f_err != nil {
		return u, err
	}

	return f_user, nil

}

func FindUser(id string) (User, error) {

	session := db.InitDb().Copy()
	defer session.Close()
	collection := db.GetCollection(session, "users")

	user := User{}
	err := collection.FindId(bson.ObjectIdHex(id)).One(&user)
	if err != nil {
		return user, err
	}

	return user, nil

}

func DeleteUser(id string) error {

	session := db.InitDb().Copy()
	defer session.Close()
	collection := db.GetCollection(session, "users")

	err := collection.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return nil

}

// Handlers

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Welcome!")
}

func ShowHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-Type", "application/json")

	if validId := bson.IsObjectIdHex(id); !validId {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(JsonHttpError{404, "Not Found"})
		return
	}

	user, err := FindUser(id)
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

	u := User{}
	u.ID = bson.NewObjectId()
	decoder := schema.NewDecoder()
	err = decoder.Decode(&u, r.Form)
	if err != nil {
		log.Fatal(err)
	}

	user, err := CreateUser(u)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UserResults{[]User{user}})

}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-Type", "application/json")

	if validId := bson.IsObjectIdHex(id); !validId {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(JsonHttpError{404, "Not Found"})
		return
	}

	uid := bson.ObjectIdHex(id)

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	session := db.InitDb().Copy()
	defer session.Close()

	collection := db.GetCollection(session, "users")

	u := &User{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(u, r.Form)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)

	err = collection.UpdateId(uid, bson.M{"$set": u})
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)

	f_user, err := FindUser(id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(UserResults{[]User{f_user}})

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-Type", "application/json")

	if validId := bson.IsObjectIdHex(id); !validId {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(JsonHttpError{404, "Not Found"})
		return
	}

	err := DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

}
