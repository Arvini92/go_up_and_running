package main

import (
	"os"
	"fmt"
	"errors"
)

func main() {
	f, err := os.Open("filename.ext")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}

	myError := errors.New("My error string")
	fmt.Println(myError)

	attendence := map[string]bool {
		"Ann": true,
		"Mike": true}

	attended, ok := attendence["Mike"]
	if ok {
		fmt.Println("Mike attended?", attended)
	} else {
		fmt.Println("No info for Mike")
	}
}
