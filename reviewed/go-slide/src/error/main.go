package main

import (
	"fmt"
	"log"
	"strconv"
)

func catcherror() {
	n, err := strconv.Atoi("hello")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("n", n)
}

func callpanic() {
	panic("Some error")

	// It didn't occur. // HL
	fmt.Println("More message") // HL
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recover from panic")
		}
	}()

	catcherror()

	callpanic()
}
