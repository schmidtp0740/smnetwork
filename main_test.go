package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestNewUser(t *testing.T) {
	message := []byte(`{
		"id": 2, 
		"firstName": "Mary", 
		"lastName": "Jane", 
		"password": "abc", 
		"dob": 1523985843719
		}`)
	request, err := http.NewRequest("POST", "/newUser", bytes.NewBuffer(message))
	if err != nil {
		panic(err)
	}
	response := httptest.NewRecorder()

	newUser(response, request)
	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, response, `{"id":2,"password":"abc","firstName":"Mary","lastName":"Jane","dob":1523985843719}`)

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkResponseBody(t *testing.T, response *httptest.ResponseRecorder, actualBody string) {
	body := response.Body.String()
	if body != actualBody {
		t.Errorf("Expected %s. Got %s", actualBody, body)

	}
}
