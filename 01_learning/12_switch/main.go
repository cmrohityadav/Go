package main

import (
	"fmt"
)


func main(){

	day:=7;

	switch day {
	case 1:
	fmt.Println("Its Mon");

	case 2:
		fmt.Println("Its Tue");
	
	case 3:
		fmt.Println("Its Wed");
	
	case 4:
		fmt.Println("Its Thu");
	
	case 5:
		fmt.Println("Its Fri");
	
	case 6:
		fmt.Println("Its Sat");
	
	case 7:
		fmt.Println("Its Sun");

	default:
		fmt.Println("Incorrect option");
	}


	num:=2;

	switch num {
	case 1:
		fmt.Println("One");
	case 2:
		fmt.Println("two");
		fallthrough
	case 3:
		fmt.Println("Three");
	case 4:
		fmt.Println("four");
	default:
		fmt.Println("Default");
	}


	// type switch

	myType:=func (i interface{}){
		switch typeOf:=i.(type){
		case int:
			fmt.Println("Int");
		case string:
			fmt.Println("String");
		case bool:
			fmt.Println("Bool");
		default:
			fmt.Println("Other",typeOf);
		}
	}

	myType(1.1);

}