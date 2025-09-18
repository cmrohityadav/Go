package main

import "fmt"

func main() {

	fmt.Println("Welcome to Pointers")

	/*
	   In C++:
	   int *ptr

	   In Go:
	   var ptr *int
	*/

	var ptr *int
	fmt.Println("Value of pointer:", ptr)

	iStockPrice := 100
	// Reference => &
	var ptrStockPrice = &iStockPrice
	fmt.Println("Address of actual pointer (StockPrice) is:", ptrStockPrice)
	fmt.Println("Value of actual pointer is:", *ptrStockPrice)

	*ptrStockPrice = *ptrStockPrice * 2
	fmt.Println("New Stock Price:", iStockPrice)

}
