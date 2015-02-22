package users_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"projects/webapi/api/users"
	"projects/webapi/db"
	"projects/webapi/routes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

// Setup

func fullUserFactory() url.Values {
	form := url.Values{}
	form.Set("first_name", "Hugo")
	form.Add("last_name", "Dorea")
	form.Add("email", "test@test.com")
	return form
}

func createTestUser(user users.User) (users.User, error) {
	session := db.InitDb().Copy()
	defer session.Close()

	collection := db.GetCollection(session, "users")

	b_user := users.User{}
	err := collection.Insert(user)
	if err != nil {
		return b_user, err
	}

	f_user := users.User{}
	err = collection.Find(bson.M{"_id": user.ID}).One(&f_user)
	if err != nil {
		return b_user, err
	}

	return f_user, nil
}

// Teardown

func cleartUserCollection() {
	session := db.InitDb().Copy()
	defer session.Close()
	db.GetCollection(session, "users").RemoveAll(bson.M{})
}

// Tests

func Test_WelComeHandler(t *testing.T) {

	db.AppEnv = "test"

	a := assert.New(t)

	m := routes.AppMux()
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/users/", nil)

	m.ServeHTTP(w, r)

	a.NoError(err)
	a.Equal(200, w.Code)
	a.Equal("application/json", w.Header().Get("Content-Type"))
	body, err := ioutil.ReadAll(w.Body)
	a.Equal(string(body), "Welcome!")

}

func Test_Show_Handler_Ok(t *testing.T) {

	db.AppEnv = "test"

	user := users.User{
		bson.NewObjectId(),
		"Hugo",
		"Dorea",
		"test@test.com",
	}

	ruser, iderr := createTestUser(user)
	if iderr != nil {
		log.Fatal(iderr)
	}
	defer cleartUserCollection()

	a := assert.New(t)

	m := routes.AppMux()
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/users/"+ruser.ID.Hex(), nil)

	m.ServeHTTP(w, r)

	a.NoError(err)
	a.Equal(200, w.Code)
	a.Equal("application/json", w.Header().Get("Content-Type"))

	body, _ := ioutil.ReadAll(w.Body)
	a.Contains(string(body), "users")
	a.Contains(string(body), "_id")
	a.Contains(string(body), ruser.ID.Hex())
	a.Contains(string(body), `"first_name":"Hugo"`)
	a.Contains(string(body), `"last_name":"Dorea"`)
	a.Contains(string(body), `"email":"test@test.com"`)

}

func Test_Show_Handler_NotFound(t *testing.T) {

	db.AppEnv = "test"

	a := assert.New(t)

	fakeId := bson.NewObjectId()

	m := routes.AppMux()
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/users/"+fakeId.Hex(), nil)

	m.ServeHTTP(w, r)

	a.Equal(404, w.Code)
	a.Equal("application/json", w.Header().Get("Content-Type"))

	body, _ := ioutil.ReadAll(w.Body)
	a.Contains(string(body), `"code":404`)
	a.Contains(string(body), `"message":"Not Found"`)

}

func Test_Create_User_Handler_Ok(t *testing.T) {

	db.AppEnv = "test"

	form := fullUserFactory()

	a := assert.New(t)

	defer cleartUserCollection()

	m := routes.AppMux()
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/users/create", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	m.ServeHTTP(w, r)

	a.NoError(err)
	a.Equal(201, w.Code)
	a.Equal("application/json", w.Header().Get("Content-Type"))

	body, _ := ioutil.ReadAll(w.Body)
	a.Contains(string(body), "users")
	a.Contains(string(body), "_id")
	a.Contains(string(body), `"first_name":"Hugo"`)
	a.Contains(string(body), `"last_name":"Dorea"`)
	a.Contains(string(body), `"email":"test@test.com"`)

}
