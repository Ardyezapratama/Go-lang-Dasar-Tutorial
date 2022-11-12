package main

import "fmt"

func main(){
	person := map[string]string{
		"name": "Eza",
		"Address": "Bali",

	}

	person["title"] = "Programmer"


	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["Address"])
	fmt.Println(person["title"])

	book := make(map[string]string)
	book["title"] = "Belajar Go-lang"
	book["author"] = "Eza Pratama"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))



/*
*	Function Map
	- len(map) Untuk mendapatkan jumlah data di map
	- map[key] Mengambil data di map dengan key
	- map[key] = value Mengubah data di map dengan key
	- make(map[TypeKey]TypeValue) Membuat map baru
	- delete(map, key) Menghapus data di map dengan key

*/
}