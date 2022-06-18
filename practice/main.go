package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	states["1"] = "1"
	states["2"] = "2"
	states["3"] = "3"
	fmt.Println(states)
	delete(states, "1")
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v : %v\n", k, v)
	}
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%v : %v\n", k, states[k])
	}
}
