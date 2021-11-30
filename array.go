package main

import "fmt"

func main(){
	var names [3]string;

	names[0] = "Raka";
	names[1] = "Janitra";
	names[2] = "Arkaan";

	fmt.Println(names[0]);
	fmt.Println(names[1]);

	var values = [3]int{
		90, 98, 87,
	};

	fmt.Println(values)
}