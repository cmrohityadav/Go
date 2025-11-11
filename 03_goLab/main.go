package main

import "fmt"

func main() {

	fmt.Println("hello world")

	var arr [4]int;

	i:=0;
	for i<len(arr){
		arr[i]=i;
		i++;
		fmt.Println("my array value at ",i,"= ",arr[i-1]);
		fmt.Printf("\n%p",&arr)
	}

}
