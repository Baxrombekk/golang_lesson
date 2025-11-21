package main

import "fmt"

type Talaba struct {
	name  string
	age   int
	slice float64
}

func main() {
	// Structdan obyekt yaratish
	var student1 Talaba
	var student2 Talaba

	student1.name = "Hege"
	student1.age = 45
	student1.slice = 60

	student2.name = "Ali"
	student2.age = 45
	student2.slice = 60

	fmt.Printf("1 %+v\n", student1)
	fmt.Printf("2 %+v\n", student2)
}
