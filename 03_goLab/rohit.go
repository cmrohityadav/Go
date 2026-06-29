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
	delete(stk,"tcs")
	var scrip =[]string{"tcs","wipro","groww","apple","sony"}

	for key,value:=range scrip{
		fmt.Println("key: ",key," Value: ",value)
		val,ok:=stk[value]
		if ok{
			fmt.Println("This  key is present",key,"Value :",val)
		}else{
			fmt.Println("This nottt key is present",key,"Value :",val)

		}
	}
}