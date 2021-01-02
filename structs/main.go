package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	p1 := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "email@email.com",
			zipCode: "12345",
		},
	}
	p1.print()
	p1p := &p1
	p1p.updateFirstName("Jimmy")
	p1.print()
	p1.updateFirstName("Jimny")
	p1.print()
}
