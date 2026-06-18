package main

import "fmt"

func main() {

	var age []int

	age = append(age, 10)
	age = append(age, 20)
	age = append(age, 30)
	age[0]=1;
	age[1]=2;
	age[2]=3;

	for i:=0;i<3;i++{
		fmt.Println(age[i])
	}

}