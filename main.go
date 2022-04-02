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
	person := person{
		firstName: "Okan",
		lastName:  "Altun",
		contactInfo: contactInfo{
			email:   "okanaltun9@gmail.com",
			zipCode: 34,
		},
	}

	personPointer := &person

	personPointer.updateName("Oki")
	fmt.Println(person)
}
func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}
