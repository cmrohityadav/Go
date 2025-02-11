package main

import "fmt"

func main() {
	fmt.Println("Welcome to Map")

	mapStockPrice := make(map[string]int)

	mapStockPrice["tcs"] = 1150
	mapStockPrice["jsw"] = 556
	mapStockPrice["adani"] = 736

	fmt.Println("List of all Stocks in mapStockPrice:", mapStockPrice)
	fmt.Println("Stock price of tcs:", mapStockPrice["tcs"])

	delete(mapStockPrice, "adani")
	fmt.Println("List of all Stocks in mapStockPrice after deletion:", mapStockPrice)

	// Looping through map
	for key, value := range mapStockPrice {
		fmt.Println(key, value)
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
