package main

import (
	"fmt"
	"time"
)

type person struct {
	Relationship []person
	FirstName    string
	LastName     string
	DateOfBirth  int64
}

func main() {
	fmt.Println("Running")
	person1 := person{
		[]person{},
		"John",
		"Doe",
		time.Now().Unix(),
	}
	fmt.Println(person1)
	person2 := person{
		[]person{},
		"Mary",
		"Jane",
		time.Now().Unix(),
	}
	person1.Relationship = append(person1.Relationship, person2)
	fmt.Println(person1)
}
