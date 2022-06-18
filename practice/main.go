package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Dates and times")
	n := time.Now()
	fmt.Println("The current time is: ", n)
	t := time.Date(2022, time.June, 18, 9, 8, 0, 0, time.UTC)
	fmt.Println("", t)
	fmt.Println(t.Format(time.ANSIC))
}
