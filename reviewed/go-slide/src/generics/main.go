package main

import "fmt"

// IFN 定義需實作的 func
type IFN interface {
	Get() int
	Set(int)
}

func genericsFN(fn IFN) {
	fn.Set(1)
	fn.Get()
}

func any(v interface{}) {
	switch val := v.(type) {
	case string:
		fmt.Println("v (string)", val)
	case int:
		fmt.Println("v (int)", val)
	}
}

func main() {}
