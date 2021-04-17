package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// jose := person{"José", "Trujillo"}
	// jose := person{firstName: "José", lastName: "Trujillo"}

	var jose person
	jose.firstName = "Jose"
	jose.lastName = "Trujillo"

	fmt.Println(jose)
	fmt.Printf("%+v\n", jose)
	fmt.Println(jose.firstName, jose.lastName)

}
