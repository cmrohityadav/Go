package main

import "fmt"

func main() {
	//LIFO :: like book stack
	defer fmt.Println("world") //imagine this line go to last
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}