package main

import "fmt"

func main(){
	var (
		nilaiAkhir = 90
		absensi = 80
	)

	fmt.Println(nilaiAkhir >= 80 && absensi >= 80)

}



/*
* Operasi && (And)
	- true && ture = true
	- true && false = false
	- false && true = false
	-false && false = false

* Operasi || (Or)
	- true || true = true
	- true  || false = true
	- false || true = true
	- false || false = false

* Operasi !
	- !true = false
	- !false = true
*/