package slices

import (
	"fmt"
	"math/rand"
	"sort"
)

type Slices struct{}

func (s Slices) Start() {
	fmt.Println("--- copyTest ---")
	copyTest()
	fmt.Println("--- sliceTest ---")
	sliceTest()
}

func copyTest() {

	length := 10
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = int(rand.Int31n(100))
	}
	sort.Ints(numbers)
	fmt.Println("Numbers:", numbers)
	worst := make([]int, 5)
	copy(worst, numbers[:5])
	fmt.Println("Worst:", worst)
}

func sliceTest() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	firstHalf := numbers[:3]
	numbers[0] = 7
	fmt.Println("Numbers:", numbers)
	fmt.Println("First half", firstHalf)
}
