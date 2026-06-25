package main

import "fmt"

func main() {

	var ltp =[4]int{0,1,2,3}
	a:=&ltp;

	fmt.Println((a))
	fmt.Println(a[2])
	fmt.Println((*a)[2])
	fmt.Println((*a))
}