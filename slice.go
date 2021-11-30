package main;

import "fmt";

func main(){
	var months = [...]string{
		"Januari", "Febuari", "Maret", "April",
		"Mei", "Juni", "Juli", "Agustus", "September",
		"Oktober", "November", "Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1);
	fmt.Println(len(slice1));
	fmt.Println(cap(slice1));

	var slice2 = months[4:6];
	fmt.Println(slice2);

	var slice3 = append(slice2, "Eko");
	fmt.Println(slice3);
	slice3[1] = "Bukan desember"
	fmt.Println(slice3);

	fmt.Println(slice2);
	fmt.Println(months);

	newSlice := make([]string, 2, 5);
	newSlice[0] = "Raka";
	newSlice[1] = "Janitra";
	newSlice[2] = "Arkaan";

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
}