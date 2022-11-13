package main

import (
	"fmt"
)

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHai(name string) {
	fmt.Println("Heloo", name, "My name is", customer.Name)

}

func main() {
	var eza Customer
	eza.Name = "Eza"
	eza.Address = "Indonesia"
	eza.Age = 22

	eza.sayHai("Ardy")

	// fmt.Println(eza)

	// ardy := Customer{
	// 	Name:    "Ardy",
	// 	Address: "Bali",
	// 	Age:     22,
	// }

	// fmt.Println(ardy)

	// pratama := Customer{"Pratama", "Banyuwangi", 22}
	// fmt.Println(pratama)
}
