package main

import "fmt"

type Man struct{
	Name string
}

func(man *Man) Married(){
	man.Name = "Mr. " + man.Name
}

func main(){
	eza := Man{"Eza"}
	eza.Married()
	fmt.Println(eza.Name)
}