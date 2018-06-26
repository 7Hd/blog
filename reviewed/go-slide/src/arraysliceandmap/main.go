package main

import (
	"fmt"
)

func main() {
	array := [3]int{1, 2, 3}
	slice1 := array[0:3]
	slice2 := []int{}
	slice3 := make([]int, 5)
	m1 := map[int]int{1: 1}
	m2 := make(map[int]int)

	slice1[0] = 4

	fmt.Println("array", array, len(array), cap(array))
	fmt.Println("slice1", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2", slice2, len(slice2), cap(slice2))
	fmt.Println("slice3", slice3, len(slice3), cap(slice3))
	fmt.Println("map1", m1, len(m1))
	fmt.Println("map2", m2, len(m2))
}
