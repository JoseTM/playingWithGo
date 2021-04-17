package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// jose := person{"José", "Trujillo"}
	// jose := person{firstName: "José", lastName: "Trujillo"}

	/* var jose person
	jose.firstName = "Jose"
	jose.lastName = "Trujillo"
	jose.contact.email = "jose.hoc@gmail.com"
	jose.contact.zip = 38007 */

	jose := person{
		firstName: "José",
		lastName:  "Trujillo",
		contact: contactInfo{
			email: "jose.hoc@gmail.com",
			zip:   38007,
		},
	}

	fmt.Println(jose)
	fmt.Printf("%+v\n", jose)
	fmt.Println(jose.firstName, jose.lastName)

}
