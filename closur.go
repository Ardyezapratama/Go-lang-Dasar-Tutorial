package main

import "fmt"

func main(){
	name := "Eza"
	counter := 0

	increment := func(){
		name = "Pratama"
		fmt.Println("Increment")
		counter++
	}

	increment()



	fmt.Println(counter)
	fmt.Println(name)
}