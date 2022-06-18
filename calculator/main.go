package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("A simple calculator")
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Value 1: ")
	val1str, _ := input.ReadString('\n')
	val1, err := strconv.ParseFloat(strings.TrimSpace(val1str), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 parsing error")
	}
	fmt.Print("Value 2: ")
	val2str, _ := input.ReadString('\n')
	val2, err := strconv.ParseFloat(strings.TrimSpace(val2str), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 parsing error")
	}
	fmt.Printf("The sum of %.2f and %.2f is %.2f\n", val1, val2, val1+val2)
}
