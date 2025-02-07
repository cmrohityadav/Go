package main

import "fmt"

// if we use first Letter Capital the it like: Public
const LoginToken string = "abcdefg"


// numberOfUser := 30000 
//above code is not allow in public,it can be only use in method / function

func main() {
	fmt.Println("Variable")

	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type :%T \n ", smallVal)

	var smallFloat float32 = 15585.123456789
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type :%T \n ", smallFloat)

	// default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n ", anotherVariable)

	// implicit type
	var website = "cmrohityadav.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n ", LoginToken)

}
