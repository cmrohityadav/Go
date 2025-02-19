package main

import (
	"fmt"
	"io"
	"net/http"
	
	"strings"
)

func main() {
	fmt.Println("Welcome to web request")
	// PerformGetRequest()
	PerformPostJsonRequest()
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

func PerformPostJsonRequest(){
	const uri="http://localhost:8000/post"

	//dummy json payload
	requestBody:=strings.NewReader(`
	{
		"courseName":"learn go",
		"price":0,
		"platform":"youtube.com"

	}
	`)
	response,err:=http.Post(uri,"application/json",requestBody)

	if err!=nil{
		panic(err)
	}

	defer response.Body.Close()
	content,_:=io.ReadAll(response.Body)
	fmt.Println(string(content))


}