package main

import "fmt"

var location string="kandivali";

var (
	city string="mumbai"
	state string="maharashtra"
	pincode int=400066

)

var fname,mname,lname string="rohit","cm","yadav"


func main() {

	fmt.Println("hello rohit");


	var name string ="rohit"
	fmt.Println(name)

	surname:="yadav"
	fmt.Println(surname)

	fmt.Println(location)
	fmt.Println("city :",city)
	fmt.Println("state :",state)
	fmt.Println("pincode :",pincode)
	fmt.Println("first name :",fname)
	fmt.Println("middle name :",mname)
	fmt.Println("last name:",lname)



}