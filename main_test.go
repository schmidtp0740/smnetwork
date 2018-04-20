package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewUser(t *testing.T) {
	fmt.Println("Testing New User")
	message := []byte(`{
		"id": 2, 
		"firstName": "Mary", 
		"lastName": "Jane", 
		"password": "abc", 
		"dob": 1523985843719
		}`)
	request, _ := http.NewRequest("POST", "/newUser", bytes.NewBuffer(message))
	response := httptest.NewRecorder()

	newUser(response, request)
	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, response, `{"id":2,"password":"abc","firstName":"Mary","lastName":"Jane","dob":1523985843719}`)

}

func TestLogin(t *testing.T) {
	fmt.Println("Testing Login")
	message := []byte(`{"id": 2, "password": "abc"}`)
	request, _ := http.NewRequest("POST", "/newUser", bytes.NewBuffer(message))
	response := httptest.NewRecorder()

	login(response, request)
	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, response, `{"response": "successful"}`)

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
