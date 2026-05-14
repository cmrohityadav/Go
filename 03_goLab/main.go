package main

import "fmt"

type Market int
const (
	EQ Market=iota+100
	F
	O
	CD

)
func main() {
	fmt.Println(F)

}
