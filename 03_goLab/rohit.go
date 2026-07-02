package main

import "fmt"

func Print(x int) {
	fmt.Println("Normal:", x)
}

func main() {
	a := 10

	defer Print(a)

	defer func() {
		fmt.Println("Closure:", a)
	}()

	defer func(x int) {
		fmt.Println("Anonymous Parameter:", x)
	}(a)

	a++

	fmt.Println("Current:", a)
}