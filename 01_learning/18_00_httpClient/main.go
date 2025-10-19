package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	//--------------simple GET Request-------------
	pHttpResponse,err:=http.Get("https://jsonplaceholder.typicode.com/posts");

	if err!=nil{
		fmt.Println("errr: ",err); 
	}
	
	defer pHttpResponse.Body.Close();

	byteResponse,err:=io.ReadAll(pHttpResponse.Body)
	if err!=nil{
		fmt.Println(string(byteResponse))
	}

	fmt.Println(string(byteResponse))



	//-------Making Request-----------------

	pHttpClient:=&http.Client{
		Timeout: time.Second*5,
	}

	pHttpRequest,err:=http.NewRequest("GET","https://jsonplaceholder.typicode.com/posts/1",nil);

	if err!=nil{
		fmt.Println("Error while making Request",err);
	}

	pHttpRequest.Header.Add("Content-Type", "application/json");

	pHttpResponse2,err:=pHttpClient.Do(pHttpRequest);
	
	if err!=nil{
		fmt.Println("Error: ",err)
	}
	
	defer pHttpResponse2.Body.Close();

	fmt.Println("RESPONSE 2 HEADER :",pHttpResponse.Header);

	byeResponse2,err:=io.ReadAll(pHttpResponse2.Body)
	if err!=nil{
		fmt.Println(err);
	}

	fmt.Println("Response Via Client making NewRequest: ",string(byeResponse2));
}