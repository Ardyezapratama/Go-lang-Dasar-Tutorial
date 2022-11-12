package main

import "fmt"

func main(){
	var (
		a = 10
		b = 20
		c = a * b
	)

	fmt.Println(c)

	var i = 10 * 10
	fmt.Println(i)

	var result = 10
	result += 10 // result = result + 10
	fmt.Println(result)

	result++ // result = result + 1
	fmt.Println(result)


	var negative = -100

	fmt.Println(negative)
}