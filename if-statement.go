package main

import "fmt"

func main(){
	name := "Ardy"

	if name == "Eza"{
		fmt.Println("Hello Eza")
	}else if name == "Pratama"{
		fmt.Println("Hello Pratama")
	}else if name == "Ardy"{
		fmt.Println("Hello Ardy")
	}else{
		fmt.Println("Hi Eza! Boleh Kenalan?")
	}


	if length := len(name); length > 5{
		fmt.Println("terlalu panjang")
	}else{
		fmt.Println("Nama Sudah Benar")
	}
}