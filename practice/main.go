package main

import (
	"fmt"
)

func main() {
	fmt.Println(addTwoValues(1, 2))
	total, cnt := addAddValues(1, 2, 3, 4, 5)
	fmt.Println(total, cnt)
}

func addTwoValues(val1 int, val2 int) int {
	return val1 + val2
}

func addAddValues(values ...int) (int, int) {
	var sum int
	for _, val := range values {
		sum += val
	}
	return sum, len(values)
}
