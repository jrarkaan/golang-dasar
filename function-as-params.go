package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string)string){
	nameFiltered := filter(name);

	fmt.Println("hello", nameFiltered);
}

func spamFilter(name string) string{
	if name == "anjing"{
		return "...";
	}else{ 
		return name;
	}
}

func main(){
	sayHelloWithFilter("raka", spamFilter)
}