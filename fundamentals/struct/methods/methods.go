package main

import (
	"fmt"
	"strings"
)

// Person ...
type Person struct {
	name    string
	surname string
}

func (person Person) getFullName() string {
	return person.name + " " + person.surname
}

func (person *Person) setFullName(fullName string) {
	splittedName := strings.Split(fullName, " ")
	person.name = splittedName[0]
	person.surname = splittedName[1]
}

func main() {
	person := Person{
		name:    "Lucas",
		surname: "Mendes",
	}

	fmt.Println(person.getFullName())

	person.setFullName("Laisla Pinto")
	fmt.Println(person.getFullName())

}
