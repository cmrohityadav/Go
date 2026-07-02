package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a/b,nil
}
func main() {
	var quotient int
	var err error

	quotient, err = Divide(4, 0)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Quotient:", quotient)
}