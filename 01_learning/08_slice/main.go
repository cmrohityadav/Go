package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var sliceStock = []string{"Tcs", "LT", "MM", "JSW"}
	fmt.Printf("Type of data of sliceStock: %T\n", sliceStock)

	sliceStock = append(sliceStock, "Bajaj", "OLA", "Maruti")
	fmt.Println(sliceStock)

	sliceStock = append(sliceStock[2:4])
	fmt.Println(sliceStock)
	fmt.Println(sliceStock[4:6])

	fmt.Println("\nNumber line from 0 to 10")
	var sliceNumberLine = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("sliceNumberLine:", sliceNumberLine)
	fmt.Println("sliceNumberLine[1:]:", sliceNumberLine[1:])
	fmt.Println("sliceNumberLine[:5]:", sliceNumberLine[:5])
	fmt.Println("sliceNumberLine[5:9]:", sliceNumberLine[5:9])

	fmt.Println("\nUsing make in slice")
	sliceHighScore := make([]int, 4);
	fmt.Println("sliceHighScore: ",sliceHighScore);
	sliceHighScore[0] = 50
	sliceHighScore[1] = 91
	sliceHighScore[2] = 52
	sliceHighScore[3] = 53
	fmt.Println(sliceHighScore)

	// sliceHighScore[4] = 533 (would cause an error)

	sliceHighScore = append(sliceHighScore, 555, 666, 777)
	fmt.Println(sliceHighScore)

	sort.Ints(sliceHighScore)
	fmt.Println(sliceHighScore)

	fmt.Println(sort.IntsAreSorted(sliceHighScore))

	fmt.Println("\n\nHow to remove a value from slice based on index")

	var sliceHoldingStocks = []string{"TATA Power", "JP Power", "OLA Electric", "Markans Pharma", "Aloc India", "JSW Infra"}
	fmt.Println(sliceHoldingStocks)

	var iSoldIndex int = 2 // Removing "OLA Electric"
	sliceHoldingStocks = append(sliceHoldingStocks[:iSoldIndex], sliceHoldingStocks[iSoldIndex+1:]...)
	fmt.Println("After selling, current holding stocks:", sliceHoldingStocks)
}
