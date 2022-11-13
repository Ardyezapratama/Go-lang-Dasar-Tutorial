package main

import "fmt"

func getFullNamed()(firstName, middleName, lastName string){
	firstName = "Ardy"
	middleName = "Eza"
	lastName = "Pratama"

	return
}


func main(){
	a, b, c := getFullNamed()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}