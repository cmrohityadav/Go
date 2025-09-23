package main

import "fmt"
func doubleNumber(num int) int{
	return 2*num;
}
func main() {
	fmt.Println("Welcome to functions")
	greeter()

	result := adder(2, 3)
	fmt.Println("Result is:", result)

	proResult:=proAdder(10,20,30)
	fmt.Println("Pro result is : ",proResult)

	iFirst,sSecond:=twoValueReturn(1,2)
	fmt.Println(iFirst,sSecond)

// take function as parameter
fmt.Println("Take return function: ",takeFunction(1,2,doubleNumber))

}

func adder(iValueOne int, iValueTwo int) int {
	return iValueOne + iValueTwo
}

func greeter() {
	fmt.Println("Namaste from Golang")
}


func proAdder(values...int)int{
	total:=0
	for _ , value:=range values{
		total+=value
	}

	return total
}

func twoValueReturn(iOne int,iTwo int)(int,string){
		iSum:=iOne+iTwo
		sMsg:="This is return msg"
		return  iSum,sMsg
}

func takeFunction(a int,b int,anyFunction func(int) int) int{

	num:=anyFunction(a);
	return num+b;

}
