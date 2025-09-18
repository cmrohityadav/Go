package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	 welcome:="welcome to user input"
	fmt.Println(welcome)

	
	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Enter stock price : ")

	
	// comma ok syntax ||  comma error syntax

	input,_:=reader.ReadString(('\n'))
	fmt.Println("Stock Ltp : ",input)

	fmt.Printf("Data type of input : %T",input)

}