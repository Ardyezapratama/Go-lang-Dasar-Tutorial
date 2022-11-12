package main

import "fmt"

func sayHello(name string, age uint8){
	fmt.Println(name, age)
}

func main(){
	sayHello("Eza", 22)
	sayHello("Pratama", 20)
}