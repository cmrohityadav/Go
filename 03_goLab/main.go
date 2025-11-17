package main

import "fmt"

type student struct {
	Name   string
	RollNo int
}

func main() {

	var rohit student
	rohit.Name = "rohit c yadav"
	rohit.RollNo = 71

	fmt.Println(rohit)

}
