package main

import "fmt"

func main(){
	name := "Pratama"

	switch name {
	case "Eza":
		fmt.Println("Hallo Eza!")
	case "Pratama":
		fmt.Println("Hallo Pratama!")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
		return
	}
	//Switch short
	/*switch length := len(name); length > 5{
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}
	*/

	length := len(name)
	switch{
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama sudah benar")
	}


}