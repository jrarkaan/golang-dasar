package main;

import "fmt";

func main(){

	var person = map[string]string{
		"name": "Eko",
		"address": "Subang",
	};

	fmt.Println(person["name"]);
	fmt.Println(person["address"]);

	var book = make(map[string]string);
	book["title"] = "Belajar Go-lang";
	book["author"] = "Eko";
}