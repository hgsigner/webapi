package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WelComeHandler(t *testing.T) {

	a := assert.New(t)

	m := userMux()
	resp := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users", nil)

	m.ServeHTTP(resp, req)

	a.NoError(err)
	a.Equal(resp.Code, 200)
	a.Equal(resp.Header().Get("Content-Type"), "application/json")
	body, err := ioutil.ReadAll(resp.Body)
	a.Equal(string(body), "Welcome!")

}

func Test_Create_User_Handler_Ok(t *testing.T) {

	a := assert.New(t)

	m := userMux()
	resp := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/users/create", strings.NewReader(`
		{
			"first_name":"Hugo", 
			"last_name":"Dorea", 
			"email":"test@test.com"
		}
		`))

	m.ServeHTTP(resp, req)

	a.NoError(err)
	a.Equal(resp.Code, 201)
	a.Equal(resp.Header().Get("Content-Type"), "application/json")
	// body, err := ioutil.ReadAll(resp.Body)
	// a.Equal(string(body), "Welcome!")

}
