package main

import "fmt"

func main(){
	months := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",

	}

	slice1 := months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//months[5] = "Diubah"
	//fmt.Println(slice1)

	//slice1[0] = "Mei Lagi"
	//fmt.Println(months)

	//Appen function
	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "Eza")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)


	//Make Slice Code
	newSlice := make([]string, 2,5)

	newSlice[0] = "Eza"
	newSlice[1] = "Pratama"

	fmt.Println(newSlice)


	//Copy Slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)


	//ini array
	iniArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)

	//ini slice
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniSlice)

}


/*
*	Tipe Data Slice
	- Tipe data slice ada 3 data, yaitu pointerr, length dan capacity
	- Pointer merupakan petunjuk data pertama di array para slice
	- Length merupakan panjang dari slice, dan
	- Capacity merupakan kapasitas dari slice, dimana length tidak boleh lebih dari capacity

*	Membuat Slice dari Array
	- array[low:high] Membuat slice dari array dimulai dari index low sampai index sebelum high
	- array[low:] Membuat slice dari array dimulai dari index low sampai index akhir dari array
	- array[:high] Membuat slice dari array dimulai dari index 0 sampai index sebelum high
	- array[:] membuat slice dari array dimulai index 0 samapi index akhir dari array

*	Function Slice
	- len(slice) Untuk mendapatkan panjang slice
	- cap(slice) Untuk mendapatkan kapasitas slice
	- append(slice, data) Membuat slice baru dengan menambahkan data ke posisi terakhir slice, jika kapasitas sudah penuh maka akan membuat array baru
	- make([]TypeData, length, capacity) Membuat slice baru
	- copy(destination, source) Menyalin slice dari source ke destination

*/