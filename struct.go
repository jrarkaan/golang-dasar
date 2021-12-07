package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

// struct function/ method
func sayHi(customer Customer, name string){
	fmt.Println("Hello", name, "My name is", customer.Name)
}

// struct function/ method
func (customer Customer) sayHi1(name string){
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main(){
	var eko Customer;
	eko.Name = "Eko";
	eko.Address = "Indonesia"
	eko.Age = 30

	sayHi(eko, "Joko");

	eko.sayHi1("Joko")

	// fmt.Println(eko.Name);

	// var raka Customer = Customer {
	// 	Name: "Raka Janitra",
	// 	Address: "Zamrud",
	// 	Age: 12,
	// }

	// fmt.Println(raka);

	// var niko Customer = Customer {
	// 	"Niko", "Regensi", 15,
	// }

	// fmt.Println(niko);

}