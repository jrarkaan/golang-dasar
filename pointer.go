package main

import "fmt"

type address struct {
	city, province, country string
}

func main(){	

	address1 := address{
		"Subang", "Jawa Barat", "Indonesia",
	}
	address2 := &address1;

	address2.city = "Bandung";

	fmt.Println(address1);
	fmt.Println(address2);
}