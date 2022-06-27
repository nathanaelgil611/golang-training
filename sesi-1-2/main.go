package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Hello, World!")

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println("Genap")
	// 	} else {
	// 		fmt.Println("Ganjil")
	// 	}
	// }

	// var user = []string{"andi", "Budi", "cacing"}
	// for _, data := range user {
	// 	fmt.Println(data)
	// }

	var student = []Biodata{
		{Name: "Andi", Age: 20, Address: "Kos", Reason: "Supaya pro"},
		{Name: "Budi", Age: 21, Address: "Kos", Reason: "Supaya pro"},
		{Name: "Cacing", Age: 22, Address: "Kos", Reason: "Supaya pro"},
	}

	fmt.Println(student[len(os.Args)-1])
}

type Biodata struct {
	Name    string
	Age     int
	Address string
	Reason  string
}
