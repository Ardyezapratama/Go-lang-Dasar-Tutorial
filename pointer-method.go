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

/* 
-	Walaupun method akan menempel di struct, tapi sebenarnya dat struct yang diakses didalam
	method adalah pass by value
-	Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena
	harus selalu diduplikasi ketika membuat method