package main

import "fmt"

func main(){
	type noKTP string
	type married bool

	var noKtpEko noKTP = "121312121"
	var marriedStatus married = true

	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
}