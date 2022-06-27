package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Genap")
		} else {
			fmt.Println("Ganjil")
		}
	}
}
