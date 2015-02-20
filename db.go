package main

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

const (
	MongoDBHosts = "127.0.0.1"
	AuthDatabase = "webapi_development"
	AuthUserName = ""
	AuthPassword = ""
	TestDatabase = "webapi_test"
)

var (
	collection = ""
)

func InitDb() *mgo.Session {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	return mongoSession

}
