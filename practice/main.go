package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 5}
	fmt.Println(poodle)
	fmt.Println(poodle.Breed)
	poodle.Weight = 9
	fmt.Println(poodle.Weight)
}

type Dog struct {
	Breed  string
	Weight int
}
