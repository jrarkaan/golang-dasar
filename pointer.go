package main

import "fmt"

type address struct {
	city, province, country string
}

func changeCountryIndonesia(address *address){
	address.country = "Indonesia";
}

func main(){	

	var address1 address = address{
		"Subang", "Jawa Barat", "Indonesia",
	}
	var address2 *address = &address1; // pointer

	address2.city = "Bandung";

	// *address2 = address{"malang", "jawa timur", "indonesia"};

	fmt.Println(address1);
	fmt.Println(address2);

	var address4 *address = new(address);
	address4.city = "Jakarta";
	fmt.Println(address4);

	var alamat = address{
		city: "subang",
		province: "jawa barat",
		country: "",
	}
	// var alamatPointer *address = &alamat
	changeCountryIndonesia(&alamat);
	fmt.Println(alamat)
}