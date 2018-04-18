package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type person struct {
	ID            int      `json:"id,omitempty"`
	Password      string   `json:"password,omitempty"`
	Relationships []person `json:"relationship,omitempty"`
	FirstName     string   `json:"firstName,omitempty"`
	LastName      string   `json:"lastName,omitempty"`
	DateOfBirth   int64    `json:"dob,omitempty"`
}

var peopleDatabase = make(map[int]person)

func main() {
	// create new mux router
	router := mux.NewRouter()
	// assign port
	port := ":8080"
	// endpoint to add new user
	router.HandleFunc("/api/signup", newUser).Methods("POST")
	// endpoint to login
	router.HandleFunc("/api/login", login).Methods("POST")
	// endpoint to add new relationship
	router.HandleFunc("/api/newRelationship", newRelationship).Methods("POST")

	fmt.Println("Listening on port" + port)
	//log http listener and serve for mux router
	log.Fatal(http.ListenAndServe(port, router))
}

// Login user to access network
func login(w http.ResponseWriter, r *http.Request) {

	var genericInterface interface{}

	reqBodyAsBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(reqBodyAsBytes, &genericInterface)
	if err != nil {
		panic(err)
	}

}

// newuser add a new person
func newUser(w http.ResponseWriter, r *http.Request) {

	// Create generic interface to insert r.Body into
	var payloadInterface interface{}
	// convert r.Body to string
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// Unmarshal the body into the generic interface
	err = json.Unmarshal(body, &payloadInterface)
	if err != nil {
		panic(err)
	}

	// Create a map variable from interface
	payloadMap := payloadInterface.(map[string]interface{})

	// conver person id to string
	personID := int(payloadMap["id"].(float64))

	// Create person struct
	personData := person{
		ID: personID,
		// new user so they have no relationsips yet
		Relationships: []person{},
		// load password from r.Body and assign as string
		Password: payloadMap["password"].(string),
		//load first name from r.Body and assign as string
		FirstName: payloadMap["firstName"].(string),
		// load last name from r.Body and assign as string
		LastName: payloadMap["lastName"].(string),
		// load dob from r.Body and assign as string
		DateOfBirth: int64(payloadMap["dob"].(float64)),
	}
	peopleDatabase[personID] = personData

	// Marshal the person object into JSON object
	personJSON, err := json.Marshal(&personData)
	if err != nil {
		panic(err)
	}

	// Declare response header as application/json
	w.Header().Set("Content-Type", "application/json")
	// Declare response status as 200
	w.WriteHeader(http.StatusOK)
	// Send response with person JSON object
	w.Write(personJSON)
}

// newRelationship - add a relationship between two people
func newRelationship(w http.ResponseWriter, r *http.Request) {

	// create generic interface for request.Body
	var genericInterface interface{}

	// convert request.Body to string
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// unmarshal body that is of type []byte into generic Interface struct
	err = json.Unmarshal(body, &genericInterface)
	if err != nil {
		panic(err)
	}

	// create body as map variable
	bodyAsMap := genericInterface.(map[string]interface{})

	personID := int(bodyAsMap["personID"].(float64))
	friendID := int(bodyAsMap["friendID"].(float64))

	person := peopleDatabase[personID]
	// had to use personTemp because of the relationship mismatch when you add friend to person before adding person to friend
	personTemp := peopleDatabase[personID]
	friend := peopleDatabase[friendID]

	person.Relationships = append(person.Relationships, friend)
	friend.Relationships = append(friend.Relationships, personTemp)
	// couldnt assign  struct in apend so this was a diversion
	peopleDatabase[personID] = person
	peopleDatabase[friendID] = friend

	peopleJSON, err := json.Marshal(peopleDatabase)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(peopleJSON)
}
