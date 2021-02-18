package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	jignesh := person{
		firstName: "Jignesh",
		lastName:  "Dhua",
		contactInfo: contactInfo{
			email:   "jignesh@email.com",
			zipCode: 4200,
		},
	}

	//fmt.Println(jignesh)

	jignesh.firstName = "Mitansh"
	jignesh.lastName = "Nirmala"

	//jignesh.print()

	jignesh.updateName("Jigs")

	jignesh.print()
	fmt.Println("*****************************")

	jigneshPointer := &jignesh
	jigneshPointer.updateLastName("Test")
	jignesh.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
	p.contactInfo.email = "jigs@email.com"
}

func (pointerToPerson *person) updateLastName(newLastName string) {
	(*pointerToPerson).lastName = newLastName
}
