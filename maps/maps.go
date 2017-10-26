package maps

import (
	"fmt"
)

type Maps struct{}

func (m Maps) Start() {
	fmt.Println("--- mapsTest ---")
	mapsTest()
}

func mapsTest() {
	words := make(map[string]int)
	words["ciao"] = 4
	words["hello"] = 5
	fmt.Println("Words:", words)
	fmt.Println("Non-existing key", words["blah"])
	delete(words, "hello")
	fmt.Println("After deletion", words)
}
