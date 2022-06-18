package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	colors = append(colors, "Purple")
	fmt.Println(colors)
	colors = append(colors[1:])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 4
	numbers[1] = 3
	numbers[2] = 2
	numbers[3] = 1
	numbers[4] = 0
	fmt.Println(numbers)
	sort.Ints(numbers)
	fmt.Println(numbers)
}
