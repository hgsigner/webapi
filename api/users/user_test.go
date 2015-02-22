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

func fullUserFactory1() url.Values {
	form := url.Values{}
	form.Set("first_name", "Hugo")
	form.Add("last_name", "Dorea")
	form.Add("email", "hugo@test.com")
	return form
}

func fullUserFactory2() url.Values {
	form := url.Values{}
	form.Set("first_name", "Flora")
	form.Add("last_name", "Dorea")
	form.Add("email", "flora@test.com")
	return form
}

func partUserFactory1() url.Values {
	form := url.Values{}
	form.Set("first_name", "Hugo")
	return form
}

func partUserFactory2() url.Values {
	form := url.Values{}
	form.Set("first_name", "Flora")
	return form
}

// Teardown

func clearUsersCollection() {
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
		"hugo@test.com",
	}

	ruser, iderr := users.CreateUser(user)
	if iderr != nil {
		log.Fatal(iderr)
	}
	defer clearUsersCollection()

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
	a.Contains(string(body), `"email":"hugo@test.com"`)

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

// test invalid id

func Test_Show_Handler_Invalid_Id(t *testing.T) {

	db.AppEnv = "test"

	a := assert.New(t)

	m := routes.AppMux()
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/users/klfjlskdjflk", nil)

	m.ServeHTTP(w, r)

	a.Equal(404, w.Code)
	a.Equal("application/json", w.Header().Get("Content-Type"))

	body, _ := ioutil.ReadAll(w.Body)
	a.Contains(string(body), `"code":404`)
	a.Contains(string(body), `"message":"Not Found"`)

}

func Test_Create_User_Handler_Ok(t *testing.T) {

	db.AppEnv = "test"

	form := fullUserFactory1()

	a := assert.New(t)

	defer clearUsersCollection()

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
	a.Contains(string(body), `"email":"hugo@test.com"`)

}

func Test_Update_User_Ok(t *testing.T) {

	db.AppEnv = "test"

	a := assert.New(t)

	user := users.User{
		bson.NewObjectId(),
		"Hugo",
		"Dorea",
		"hugo@test.com",
	}

	ruser, iderr := users.CreateUser(user)
	if iderr != nil {
		log.Fatal(iderr)
	}
	defer clearUsersCollection()

	form := partUserFactory2()

	m := routes.AppMux()
	w := httptest.NewRecorder()
	r, err := http.NewRequest("PUT", "/users/"+ruser.ID.Hex(), strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	m.ServeHTTP(w, r)

	a.NoError(err)
	a.Equal(200, w.Code)
	a.Equal("application/json", w.Header().Get("Content-Type"))

	body, _ := ioutil.ReadAll(w.Body)
	a.Contains(string(body), "Flora")
	a.Contains(string(body), "Dorea")
	a.Contains(string(body), "hugo@test.com")

}

func Test_Delete_User_Ok(t *testing.T) {

	db.AppEnv = "test"

	a := assert.New(t)

	user := users.User{
		bson.NewObjectId(),
		"Hugo",
		"Dorea",
		"hugo@test.com",
	}

	ruser, iderr := users.CreateUser(user)
	if iderr != nil {
		log.Fatal(iderr)
	}
	defer clearUsersCollection()

	m := routes.AppMux()
	w := httptest.NewRecorder()
	r, err := http.NewRequest("DELETE", "/users/"+ruser.ID.Hex(), nil)

	m.ServeHTTP(w, r)

	a.NoError(err)
	a.Equal(200, w.Code)
	a.Equal("application/json", w.Header().Get("Content-Type"))

}
