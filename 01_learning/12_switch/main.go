package main

import "fmt"

func main() {
	fmt.Println("Welcome to Conditions: if else")

	var iLoginCount int = 10
	var sMessage string

	if iLoginCount < 10 {
		sMessage = "Regular User"
	} else if iLoginCount > 10 {
		sMessage = "Watch Out"
	} else {
		sMessage = "Exactly 10 login count"
	}

	fmt.Println(sMessage)

	fmt.Println("\n\nSomething new in Go")
	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	fmt.Println("\n\nIMP for Web Request Handling")

	if iNum := 3; iNum < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than or equal to 10")
	}

	// Example usage in error handling
	// if err != nil {
	//     fmt.Println("An error occurred")
	// }
}
