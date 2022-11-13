package main

import (
	"fmt"
)

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var eza Customer
	eza.Name = "Eza"
	eza.Address = "Indonesia"
	eza.Age = 22

	fmt.Println(eza)

	ardy := Customer{
		Name:    "Ardy",
		Address: "Bali",
		Age:     22,
	}

	fmt.Println(ardy)

	pratama := Customer{"Pratama", "Banyuwangi", 22}
	fmt.Println(pratama)
}
