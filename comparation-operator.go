package main

import "fmt"

func main(){
	var (
		name1 = "EZA"
		name2 = "EZA"
		value1 = 100
		value2 = 200
	)


	var result bool = name1 == name2
	fmt.Println(result)


	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)


}