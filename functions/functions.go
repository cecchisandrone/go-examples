package functions

import (
	"fmt"
)

type Functions struct{}

func (f Functions) Start() {
	fmt.Println("--- funcTest ---")
	funcTest(3, 5)
}

type adder func(a int, b int) int

func funcTest(a int, b int) {
	fmt.Println("Sum is:", calc(func(a int, b int) int {
		return a + b
	}, a, b))
}

func calc(adder adder, a int, b int) int {
	return adder(a, b)
}
