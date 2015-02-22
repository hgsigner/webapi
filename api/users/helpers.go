package users

import (
	"projects/webapi/db"

	"gopkg.in/mgo.v2/bson"
)

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
