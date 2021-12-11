package main

import "fmt"

func getHello(name string)string{
	if name == ""{
		return "Hello, bro!";
	}else{
		return "Hello, " + name;
	}
}

func getFullName()(string, string, string){
	return "Eko", "Raka", "Janitra"
}

func main(){
	result := getHello("eko");

	fmt.Println(result);

	firstName, _, _ := getFullName();

	fmt.Println(firstName);
}