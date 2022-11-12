package main

import "fmt"

func main(){
	/*counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	*/

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke ", i)
		
	}

	slice := []string{"Ardy", "Eza", "Pratama", "Junior", "Programmer"}
	for x := 0; x < len(slice); x++{
		fmt.Println(slice[x])
	}

	for x, value := range slice{
		fmt.Println("Index", x, "=", value)
	}


	person := make(map[string]string)
	person["name"] = "Eza"
	person["title"] = "Programmer"

	for key, value := range person{
		fmt.Println(key, "=", value)
	}

}