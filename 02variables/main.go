package main

import "fmt"

// In golang we use initially Capital later to declare public variable/method/function.
const LoginToken string = "ghabbhhjd" // Public

func main() {
	var username string = "Rishi"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	// lexer will automatically decide variable type for you here.
	var website = "learncodeonline.in"
	fmt.Println(website)

	// // We can't do the following later onwards.
	// website = 18

	// no var style
	// Inside a method like main and etc we are allowed to use walrus operator.
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
