package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestIndex(t *testing.T) {
	req1, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr1 := newRequestRecorder(req1, "GET", "/", Index)
	if rr1.Code != 200 {
		t.Error("Expected response code to be 200")
	}
}

func createDummyUser() *Profile {
	dummyProfile := Profile{
		user: User{
			name:     "test",
			password: "something",
			email:    "aa@gmail.com",
		},
		firstName: "test",
		lastName:  "last",
		role: Role{
			name: "user",
		},
	}
	return &dummyProfile
}

func TestGetUsername(t *testing.T) {
	user := createDummyUser()
	expected := "test"
	got := user.getUsername()
	if got != expected {
		t.Error("Found wrong Username")
	}
}

func TestCreateUser(t *testing.T) {
	user := createDummyUser()
	data, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("POST", "/user", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	requestChecker := newRequestRecorder(request, "POST", "/user", CreateUser)
	if requestChecker.Code != http.StatusCreated {
		t.Error("Create user did not successfully save")
	}
}

func newRequestRecorder(req *http.Request, method string, strPath string, fnHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, strPath, fnHandler)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}
