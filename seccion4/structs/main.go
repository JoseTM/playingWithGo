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

	jose.print()
	jose.updateName("Jose")
	jose.print()
	jose.updateNamePointer("Jose")
	jose.print()

	josePointer := &jose //un puntero a la variable
	joseSafe := jose

	josePointer.updateNameByPointer("Josith")

	jose.print()
	joseSafe.print()

	joseSafe.updateNameByPointer("patata")
	joseSafe.print()
	jose.print()
}

func (p *person) updateNamePointer(name string) {
	p.firstName = name
}

func (p person) updateName(name string) {
	p.firstName = name
}

func (p *person) updateNameByPointer(name string) {
	(*p).firstName = name
	//p.firstName = name
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
