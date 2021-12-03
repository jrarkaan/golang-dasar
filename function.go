package main

import "fmt"

func sayHello(firstname string, lastname string){
	fmt.Println("Hello,", firstname, lastname)
}

func main(){
	var firstname string = "Raka";
	sayHello(firstname, "Janitra");
}