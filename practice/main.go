package main

import (
	"fmt"
)

func main() {
	a := 42
	p := &a
	fmt.Println("Value of p:", *p)

}
