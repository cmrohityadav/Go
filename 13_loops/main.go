package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")

	days := []string{"sun", "tue", "wed", "fri", "sat"}
	fmt.Println(days)

	fmt.Println("\n\nUsing for loop")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("\n\nUsing range")
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("\n\nUsing for-each loop with range")
	for index, value := range days {
		fmt.Printf("Index is %v and value is %v\n", index, value)
	}

	fmt.Println("\n\nUsing for-each loop with range (ignoring index)")
	for _, value := range days {
		fmt.Printf("Value is %v\n", value)
	}

	fmt.Println("\n\nWhile loop, break and continue")

	rougueValue := 1
	iBreak := 1000
	iContinue := 1000

	// While loop equivalent in Go
	for rougueValue < 10 {
		if rougueValue == iContinue {
			rougueValue++
			continue
		}
		if rougueValue == iBreak {
			break
		}
		if rougueValue == 2 {
			goto cmrohit
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

	fmt.Println("\n\nGoto statement example")
	// Label for goto
cmrohit:
	fmt.Println("Jumping to github.com/cmrohityadav")
	fmt.Println("Program ending")
}
