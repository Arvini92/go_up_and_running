package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors) -1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 1
	numbers[1] = 6
	numbers[2] = 7
	numbers[3] = 10
	numbers[4] = 9
	fmt.Println(numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	sort.Ints(numbers)
	fmt.Println(numbers)
}
