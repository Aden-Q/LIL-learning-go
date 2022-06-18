package main

import (
	"fmt"
)

func main() {
	if x := -42; x < 0 {
		fmt.Println("Less than 0")
	} else if x == 0 {
		fmt.Println("Equal to 0")
	} else {
		fmt.Println("Greater than 0")
	}
}
