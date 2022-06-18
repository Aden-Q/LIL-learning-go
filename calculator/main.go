package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go"
	file, err := os.Create("sample.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote %v characters to file %v\n", length, "sample.txt")
	defer file.Close()
	defer readFile("sample.txt")
}

func readFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	checkError(err)
	fmt.Println("Read from file:", string(content))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
