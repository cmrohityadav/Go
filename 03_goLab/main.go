package main

import "fmt"

type student struct {
	Name   string
	RollNo int
}

func main() {

	p:=student{"Rohit",40}
	

	q:=&p;

	fmt.Println(q.Name,q.RollNo)

}
