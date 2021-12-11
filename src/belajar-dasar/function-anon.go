package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blackList Blacklist){
	if blackList(name){
		fmt.Println("you are blocked!", name);
	}else{
		fmt.Println("Welcome", name);
	}
}


func main(){
	blackList := func(name string)bool {
		return name == "Admin";
	}

	registerUser("Admin", blackList)
	registerUser("raka", blackList)
}