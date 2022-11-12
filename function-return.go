package main

import "fmt"

func getHello(name string) string{
	if name == ""{
		return "Hello Brodieeeee"

	}else{
		return "Hello " + name
	}
}


func main(){
	result := getHello("Eza")
	fmt.Println(result)

	fmt.Println(getHello(""))
}