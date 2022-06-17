package main

import (
	"fmt"
)

const a int = 64

func main() {
	str := "This is Go!"
	fmt.Println(str)
	fmt.Printf("The variable's type is %T\n", str)
	fmt.Println(a)
}
