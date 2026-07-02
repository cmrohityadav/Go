package main

import "fmt"

func Sum() (result int) {

	defer func() {
		result=10
		fmt.Println("defer",result)
	}()

	return result
}

func main() {
	fmt.Println(Sum())
}