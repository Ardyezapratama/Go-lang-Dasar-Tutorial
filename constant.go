package main

import "fmt"

func main() {
	const (
		firstName string = "Ardy Eza"
		lastName         = "Pratama"
		value            = 1000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}

/*
* Constant
	- Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
	- Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, buka var
	- Saat pembuatan constant, wajib langsung menginisialisasikan datanya
*/
