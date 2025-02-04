package main

import "fmt"

func main() {
	fmt.Println("Variable")


	var username string="hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n",username)


	var isLoggedIn bool =true;
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n",isLoggedIn)


	var smallVal uint8=255;
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type :%T \n ",smallVal);


	var smallFloat float32=15585.123456789;
	fmt.Println(smallFloat);
	fmt.Printf("Variable is of type :%T \n ",smallFloat);


	// default values and aliases
	var anotherVariable int;
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n ",anotherVariable)

	
}