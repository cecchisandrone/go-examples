package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type Concurrency struct{}

var (
	counter = 0
	lock    sync.Mutex
)

func (c Concurrency) Start() {

	fmt.Println("--- locksTest ---")
	locksTest()
	fmt.Println("--- channelsTest ---")
	channelsTest()
}

func locksTest() {
	for i := 0; i < 20; i++ {
		go incr()
	}
	time.Sleep(time.Second * 1)
	fmt.Println("Counter is", counter)
}

func incr() {
	lock.Lock()
	defer lock.Unlock()
	counter++
}

func channelsTest() {
	c := make(chan int)
	numbers := []int{1, 2, 3, 4, 5}

	square := func(number int) {
		square := number * number
		fmt.Print(square, " ")
		time.Sleep(1 * time.Second)
		c <- square
	}

	fmt.Println("Processing started at", time.Now())
	for i := 0; i < len(numbers); i++ {
		go square(numbers[i])
	}

	for i := 0; i < len(numbers); i++ {
		<-c
	}

	fmt.Println("\nProcessing ended at", time.Now())
}
