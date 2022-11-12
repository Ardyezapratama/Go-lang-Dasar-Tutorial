package main

import "fmt"

func main(){
	type(
		noHp string
		age uint8
	)

	var noHpEza noHp = "03482058595"
	var ageEza age = 22
	fmt.Println(noHpEza)
	fmt.Println(ageEza)

}