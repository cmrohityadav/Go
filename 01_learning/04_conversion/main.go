package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to Stock App")

	fmt.Println("Pls rate our Stock app between 1 and 5: ")
	rating := bufio.NewReader(os.Stdin)
	ratingString, _ := rating.ReadString('\n')
	fmt.Println("Thanks for rating: ", ratingString, "star")

	fRating, err := strconv.ParseFloat(strings.TrimSpace(ratingString), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating:", fRating+1)
	}
}
