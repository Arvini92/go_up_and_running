package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Close the file!")
	fmt.Println("Open the file!")
	myFunc()
	defer fmt.Println("Statment 1")
	defer fmt.Println("Statment 2")
	defer fmt.Println("Statment 3")
	defer fmt.Println("Statment 4")
	fmt.Println("Undeferred statement")

	x := 1000
	defer fmt.Println("Value of x:", x)
	x++
	fmt.Println("Value of x after incrementing:", x)
}

func myFunc() {
	defer fmt.Println("Deferred in the function")
	fmt.Println("Not deferred in the function")
}
