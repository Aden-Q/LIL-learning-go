package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var b float64
	sum := math.Round(float64(a) + b)
	fmt.Println(sum)
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum is:", intSum)
	f1, f2, f3 := 12.1, 23.5, 43.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum is:", floatSum)
	floatSum = math.Round(floatSum)
	fmt.Println("Round sum:", floatSum)

	circleRadius := 12.32323232
	fmt.Printf("Circle radius: %.2f\n", circleRadius)
}
