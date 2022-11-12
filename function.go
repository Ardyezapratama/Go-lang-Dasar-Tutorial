package main

import "fmt"

func sayHello(){
	fmt.Println("Hello Eza")

}

func main(){
	for i := 0; i < 10; i++ {
		sayHello()
	}
}