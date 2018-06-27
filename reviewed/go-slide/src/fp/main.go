package main

import (
	"fmt"
)

func hrfn(fn func(int) int) int {
	return fn(0)
}

func closure() func() int {
	n := -1

	return func() int {
		n++
		return n
	}
}

func main() {
	fn := func(i int) int { return i + 1 }
	fmt.Println("hrfn(fn)", hrfn(fn))
	fnc := closure()
	fmt.Println("fnc()", fnc())
	fmt.Println("fnc()", fnc())
}
