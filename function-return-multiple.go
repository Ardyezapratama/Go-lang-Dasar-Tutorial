package main

import "fmt"

func getFullName()(string, string, string){
	return "Ardy", " Eza", " Pratama"

}

func main(){
	_, middleName, lastName := getFullName()
	fmt.Print(middleName, lastName)
	
}