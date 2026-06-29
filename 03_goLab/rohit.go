package main

import "fmt"

func main() {

	stk:=map[string]int{
		"tcs":2400,
		"wipro":1400,
		"groww":205,
		"apple":5000,
	}

	fmt.Println(stk)

	for key,value:=range stk{
		fmt.Println("key: ",key," Value: ",value)
	}
}