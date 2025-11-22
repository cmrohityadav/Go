package main

import "errors"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can not divided by zero")
	}

	return a/b, nil;

}
func main() {

	result,err:=divide(10,0);
	if err!=nil{
		println("Error: ",err.Error())
	}

	println("Result: ",result)
}