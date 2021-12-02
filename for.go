package main
import "fmt"

func main(){

	// var counter int8
	// for counter = 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke-", counter);
	// }

	slice := []string{"eko", "kurniawan", "khannedy"}

	fmt.Println(len(slice));
	
	for i := 0; i <= len(slice) - 1; i++{
		fmt.Println(slice[i]);
	}

}