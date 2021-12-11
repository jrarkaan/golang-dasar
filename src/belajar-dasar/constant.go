package main

import "fmt"

func main(){
	const firstName string = "Eko"
	const lastName = "Raka"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		a string = "raka"
		b = "a"
		c uint8 = 100
	)
}