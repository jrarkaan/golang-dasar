package main

import "fmt"

func getFullName1()(firstName string, middleName string, lastName string){
	firstName = "Raka"
	middleName = "Janitra"
	lastName = "Arkaan"

	return
}

func main(){
	firstName, middleName, _ := getFullName1()

	fmt.Println(firstName)
	fmt.Println(middleName)
}