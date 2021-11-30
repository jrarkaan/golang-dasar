package main

import "fmt"

func main(){
	var name string 

	name = "Eko kurniawan"
	fmt.Println(name)

	name = "Raka janitra"
	fmt.Println(name)

	var age int8 = 30
	fmt.Println(age)

	country := "Indonesia" // without key var
	fmt.Println(country)

	var (
		firstName = "Eko"
		lastName = "Raka"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}