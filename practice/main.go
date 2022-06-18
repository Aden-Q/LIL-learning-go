package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	fmt.Println("Dog weight:", poodle.getWeight())
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
}

func (d Dog) getWeight() int {
	return d.Weight
}
