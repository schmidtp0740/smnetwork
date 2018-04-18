package main_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestNewUser(t *testing.T) {
	message := []byte(`{
		"id": 123,
		"firstName": "John",
		"lastName": "Doe",
		"dob": 1523985843719
	}`)
	req, _ := http.NewRequest("POST", "/newUser", bytes.NewBuffer(message))
	response := httptest.NewRecorder()
	mux.NewRouter().ServeHTTP(response, req)

	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, response)

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

func checkResponseBody(t *testing.T, response *httptest.ResponseRecorder) {
	body := response.Body.String()
	if body == `{"id": 123, firstName": "John","lastName": "Doe",	"dob": 1523985843719}` {
		t.Errorf("Received an empty array. Got %s", body)

	}
	fmt.Println(response.Body.String())
}
