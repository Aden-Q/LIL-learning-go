package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	contents := string(bytes)
	fmt.Print(contents)
}
