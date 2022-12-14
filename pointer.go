package main
import "fmt"

type Address struct{
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address){
	address.Country = "Indonesia"

}

func main(){
	address1 := Address{"Banyuwangi","Jawa Timur", "Indonesia"}
	// address2 := address1 // pointer pass by value
	address2 := &address1 // pointer pass by reference
	var address3 = &address1

	address2.City = "Jember"


	*address2 = Address{"Badung", "Bali", "Indonesia"}


	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	address4.City = "Denpasar"
	fmt.Println(address4)

	alamat := &Address{
		City: "Subang",
		Province: "Jawa Barat",
		Country: "",

	}
	ChangeCountryToIndonesia(alamat)
	fmt.Println(alamat)
}