package main

import "fmt"

func random() interface{} {
	return "eza"
}

func main() {
	// result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) // panix
	// fmt.Println(resultInt)

	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	case bool:
		fmt.Println("Boolean", value)
	default:
		fmt.Println("Unknown")
	}
}
