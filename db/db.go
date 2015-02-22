package db

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

const (
	mongoDBHosts       = "127.0.0.1"
	devDatabase        = "webapi_development"
	testDatabase       = "webapi_test"
	productionDatabase = "webapi"
	authUserName       = ""
	authPassword       = ""
)

var (
	AppEnv    = "development"
	currentDb = ""
)

func InitDb() *mgo.Session {

	switch AppEnv {
	case "development":
		currentDb = devDatabase
	case "test":
		currentDb = testDatabase
	case "production":
		currentDb = productionDatabase
	}

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{mongoDBHosts},
		Timeout:  60 * time.Second,
		Database: currentDb,
		Username: authUserName,
		Password: authPassword,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	return mongoSession

}

func GetCollection(session *mgo.Session, collection string) *mgo.Collection {
	return session.DB(currentDb).C(collection)
}
