package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	jose := person{
		firstName: "José",
		lastName:  "Trujillo",
		contactInfo: contactInfo{
			email: "jose.hoc@gmail.com",
			zip:   38007,
		},
	}

	fmt.Println(jose)
	fmt.Printf("%+v\n", jose)
	fmt.Println(jose.firstName, jose.lastName)

}
