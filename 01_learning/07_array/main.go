package main

import "fmt"

func main() {

	fmt.Println("Welcome to Array")

	// Define size (mandatory in Go)
	var arrayStocks [6]string
	arrayStocks[0] = "TATA"
	arrayStocks[1] = "Wipro"
	arrayStocks[2] = "L&T"
	arrayStocks[3] = "Infosys"
	// If a value is not assigned at an index, it remains empty but contributes to the size
	arrayStocks[5] = "Adani"

	fmt.Println("List of Stocks:", arrayStocks)
	fmt.Println("Number of Stocks:", len(arrayStocks))

	var arrayIndex = [10]string{"Nifty50", "BankNifty", "Sensex"}
	fmt.Println("Top 3 indices of India:", arrayIndex)
	fmt.Println("Size of arrayIndex:", len(arrayIndex))


	var vals[4]bool;
	vals[2]=true;
	fmt.Println(vals);

	

}
