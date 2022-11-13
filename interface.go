package main

import "fmt"

type HasName interface {
	Getname() string
}

func SayHelloGuys(hasName HasName) {
	fmt.Println("Hello", hasName.Getname())
}

type Person struct {
	Name string
}

func (person Person) Getname() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) Getname() string {
	return animal.Name
}

func main() {
	var eza Person
	eza.Name = "Eza"

	SayHelloGuys(eza)

	var dog Animal
	dog.Name = "Anjing"

	SayHelloGuys(dog)

}
