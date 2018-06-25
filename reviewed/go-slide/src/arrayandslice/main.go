package main

import (
	"fmt"
)

func main() {
	array := [3]int{1, 2, 3}
	slice := array[0:3]
	slice5 := make([]int, 5)

	slice[0] = 4

	fmt.Println("array", array)
	fmt.Println("slice", slice)
	fmt.Println("slice5", slice5)
}
