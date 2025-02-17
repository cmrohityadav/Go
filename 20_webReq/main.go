package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web request")
	PerformGetRequest()
}

func PerformGetRequest(){

	const myuri="http://localhost:8000/get"
	response,err:=http.Get(myuri)

	if err!=nil{
		panic(err)
	}


	defer response.Body.Close()

	fmt.Println("Status code: ",response.StatusCode)
	fmt.Println("Content length is : ",response.ContentLength)
	contentInBytes,_:=io.ReadAll(response.Body)
	fmt.Println(string(contentInBytes))

	fmt.Println("Using Strings")

	var responseString strings.Builder
	byteCount,_:=responseString.Write(contentInBytes)
	fmt.Println("byte count : ",byteCount)
	fmt.Println(responseString.String())

}