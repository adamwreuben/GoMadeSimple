package main

import (
	"fmt"
	"strconv"
)

//Define person struct

//Methods in struc
// value receiver and pointer receiver

type Person struct {
	firstname string
	lastname  string
	city      string
	gender    string
	age       int
}

//Methods for person go here ---> THIS IS VALUE RECEIVE METHOD
func (p Person) greet() string {
	return "Hello, my name is " + p.firstname + " am " + strconv.Itoa(p.age) //p is like this
}

//hasBirthday method (POINTER RECEIVER)
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried () Pointer receiver
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastname = spouseLastName
	}
}

func main() {

	//init person using struct
	person1 :=
		Person{firstname: "Adam", lastname: "Reuben", age: 23, city: "Dar", gender: "M"}

	person2 :=
		Person{firstname: "Eliza", lastname: "Reuben", age: 23, city: "Dar", gender: "F"}
	//alternative pass as position
	// person2 :=
	// Person{ "Adam",  "Reuben",23,  "Dar",  "M"}

	//person1.firstname
	person1.hasBirthday()
	person2.getMarried("Reuben")
	person1.getMarried("Williams")
	fmt.Print(person1.greet())

}
