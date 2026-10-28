package main

import "fmt"

type NameProvider interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Company struct {
	Name string
}

func (company Company) GetName() string{
	return company.Name
}

func SayHello(value NameProvider) {
	fmt.Println("Hello, ", value.GetName())
}
func main() {
	person := Person{Name: "Alice"}
	company := Company{Name: "Acme Inc."}
	SayHello(person)
	SayHello(company)
}