package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	person := person{
		firstName: "Okan",
		lastName:  "Altun",
		contact: contactInfo{
			email:   "okanaltun9@gmail.com",
			zipCode: 34,
		},
	}
	fmt.Println(person.contact.email)
}
