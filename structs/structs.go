package structs

import "fmt"

// User type
type User struct {
	FirstName string
	LastName  string
}

func (u *User) change(name string) {
	u.FirstName = name
}

func Start() {
	fmt.Println("--- pointerTest ---")
	pointerTest()
}

func pointerTest() {
	alex := &User{"Alessandro", "Dionisi"}

	alex.change("Alex")
	fmt.Printf("Pointer: %p\n", alex)
	fmt.Println("Alex modified:", *alex)
}
