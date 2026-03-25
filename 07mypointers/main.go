package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)

}

// &: When to use it: When you need to pass a reference to a variable rather than a copy of its value.
// *:
// A. When * is in front of a type (like *int or *User), it’s just a label. It tells Go, "This variable doesn't hold data; it holds a memory address."
// B. When you put * in front of a variable that is already a pointer, you are dereferencing it. It means, "Follow the map and show me what’s actually inside the house."
