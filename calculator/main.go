package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Value 1: ")
	float1 := readFloat(reader)

	fmt.Print("Value 2: ")
	float2 := readFloat(reader)

	fmt.Print("Select an operation (+ - * /):")
	input3, _ := reader.ReadString('\n')
	op := strings.TrimSpace(input3)
	result := calTwoValues(float1, float2, op)
	fmt.Printf("The result is %f\n", result)
}

func readFloat(reader *bufio.Reader) float64 {
	input, _ := reader.ReadString('\n')
	val, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}
	return val
}

func calTwoValues(val1, val2 float64, op string) float64 {
	var result float64
	switch op {
	case "+":
		result = val1 + val2
	case "-":
		result = val1 - val2
	case "*":
		result = val1 * val2
	case "/":
		result = val1 / val2
	}
	return result
}
