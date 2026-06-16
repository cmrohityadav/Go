package main

import "fmt"

func main() {

	isAdmin:=true
	isLogin:=true
	age:=19

	if isAdmin && isLogin && age>18{
		fmt.Println("Welcome")
	}else{
		fmt.Println("Welcome Not")
	}

	perItem:=10;
	item:=5;
	if total:=item*perItem;total>100{
		fmt.Println("Free Delivery Avaialble")
	}else{
		fmt.Println("NO Free Delivery Avaialble")
	}


}