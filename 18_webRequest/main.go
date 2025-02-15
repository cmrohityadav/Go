package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://cmrohityadav.github.io/projects/12_08Mar23_lightOnOff/index.html"

func main() {
	fmt.Println("welcome to web request")

	res,err:=http.Get(url)

	if err!=nil{
		panic(err)
	}

	fmt.Printf("Response Type : %T \n",res)
 defer	res.Body.Close()//caller's responsibility to close connection

 dataBytes,err:=ioutil.ReadAll(res.Body)
	
 if err!=nil{
	panic(err)
}

content:=string(dataBytes)
fmt.Println("content  from response : ",content)




}