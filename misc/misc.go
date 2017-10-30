package misc

import "fmt"

type Misc struct{}

func (s Misc) Start() {
	fmt.Println("--- deferTest ---")
	deferTest()
	fmt.Println("--- conversionTest ---")
	conversionTest()
}

func deferTest() {
	fmt.Println("Entering deferTest")
	defer fmt.Println("After exiting deferTest")
	fmt.Println("Exiting deferTest")
}

func conversionTest() {
	genericFunc(9)
}

func genericFunc(num1 interface{}) {

	switch num1.(type) {
	case int:
		fmt.Println("name is an int")
	case string:
		fmt.Println("name is a string")
	}

	num2 := num1.(int)
	fmt.Println("num1,num2:", num1, num2)
}
