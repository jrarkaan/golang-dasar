package main
import "fmt"

func main(){

	// var counter int8
	// for counter = 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke-", counter);
	// }
	
	// for slice
	slice := []string{"eko", "kurniawan", "khannedy"}

	fmt.Println(len(slice));
	
	for i := 0; i <= len(slice) - 1; i++{
		fmt.Println(slice[i]);
	}

	for i, value := range slice {
		fmt.Println("Index-", i, "=", value);
	}

	// for map
	person := make(map[string]string);

	person["name"] = "Eko"
	person["title"] = "Raka"

	for key, value := range person {
		fmt.Println(key, "=", value);
	}
	
}