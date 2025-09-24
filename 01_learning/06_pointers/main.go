package main

import "fmt"

func changeByValue(num int){
	num++;
	fmt.Println("The value num in changeByValue(): ",num);
}

func changeByReference(num *int){
	(*num)++;

	fmt.Println("The value num in changeByReference(): ",*num);

}

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

	var num int=10;
	fmt.Println("\nValue of num in main func Before calling changeByValue(): ",num);
	changeByValue(num);
	fmt.Println("Value of num in main func after called changeByValue(): ",num);


	fmt.Println("\nValue of num in main func Before calling changeByReference(): ",num);
	changeByReference(&num);
	fmt.Println("\nValue of num in main func after called changeByReference(): ",num);




}
