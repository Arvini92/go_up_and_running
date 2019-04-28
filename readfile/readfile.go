package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := "./fromString.txt"

	content, err := ioutil.ReadFile(fileName)
	checkError(err)

	result := string(content)

	fmt.Println("Read from file:", result)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
