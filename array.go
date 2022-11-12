package main

import "fmt"

func main(){
	var names = [3]string{
		"Ardy", "Eza", "Pratama",
	}

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3] int{
		90, 80, 70,
	}

	var lagi[10] string


	values[2] = 100
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))
	fmt.Println(len(lagi))

}

/*
*	Function Array
	- len(array)  Untuk mendapatkan panjang array
	- array[index]  Mendapat data di posisi index
	- array[index] = value Untuk mengubah data posisi index

*/