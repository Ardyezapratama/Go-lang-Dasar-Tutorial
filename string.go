package main

import "fmt"

func main() {
	fmt.Println("Ardy")
	fmt.Println("Ardy Eza")
	fmt.Println("Ardy Eza Pratama")

	fmt.Println(len("Ardy"))
	fmt.Println("Ardy Eza"[0])
	fmt.Println("Ardy Eza Pratama"[2])

}

/*
* Tipe Data String
	- String merupakan data kumpulan karakter
	- Jumlah karkater di dalam String bisa nol sampai tidak terhingga
	- Tipe data String di Go-lang drepresentasikan dengan kata string
	- Nilai dat String di Go-lang selalu diawali dengan karakter "(petik dua) dan diakhiri dengan karakter "(petik dua)

*Function untuk String
	- len("string") untuk menghitung jumlah karakter di String
	- "string[number]" untuk mengambil karakter pada posisi yang ditentukan
*/
