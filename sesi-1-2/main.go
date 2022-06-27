package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println("Genap")
	// 	} else {
	// 		fmt.Println("Ganjil")
	// 	}
	// }

	var user = []string{"andi", "Budi", "cacing"}
	for _, data := range user {
		fmt.Println(data)
	}
}
