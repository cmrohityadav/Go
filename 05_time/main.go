package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	/* refer and learn it
		01 => Month
		02 => Day
		2006 => Year
		Monday => Present week day
		15 => Hour
		04 => Minute
		05 => Second
	*/
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("02-01-2006"))

	fmt.Println("\nADDING DAY OF WEEK")
	fmt.Println(presentTime.Format("02-01-2006 Monday"))
	fmt.Println(presentTime.Format("Monday 02-01-2006"))
	fmt.Println(presentTime.Format("02 Monday 01-2006"))
	fmt.Println(presentTime.Format("02-Monday-01-2006"))

	fmt.Println("\nADDING TIME")
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	fmt.Println("\nCREATING DATE")
	createdDate := time.Date(2000, time.June, 23, 9, 30, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006-Monday-15-04-05"))
}
