package main

import (
	"fmt"
	"net/url"
)

const myurl string="https://www.udemy.com/join/passwordless-auth/?next=%2Fhome%2Fmy-courses%2Flearning%2F"

func main() {
	fmt.Println("Welcome to handling urls in golang")

	fmt.Println(myurl)

	//parsing
	result,_:=url.Parse(myurl)
	fmt.Println("Result :",result)
	fmt.Println("Scheme :",result.Scheme)
	fmt.Println("Host :",result.Host)
	fmt.Println("Path :",result.Path)
	fmt.Println("Port :",result.Port())
	fmt.Println("RawQuery :",result.RawQuery)


	qparams:=result.Query()
	fmt.Printf("The type of query params are: %T\n",qparams)
	fmt.Println(qparams["next"])

	partOfUrl:=&url.URL{
		Scheme: "https",
		Host: "cmrohityadav.in",
		Path: "/about",
		RawPath: "user=rohit",
	}
	anotherPath:=partOfUrl.String()
	fmt.Println("anotherpath : ",anotherPath)

}