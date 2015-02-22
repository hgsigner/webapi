package users

import "gopkg.in/mgo.v2/bson"

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
