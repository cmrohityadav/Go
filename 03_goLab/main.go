package main

import (
	"fmt"
	"net/http"
)

func anyHandleFunc(res http.ResponseWriter,req *http.Request){
	fmt.Fprintln(res,"Hello rohit")	
}
func main() {
	fmt.Println("Welcome to HTTP")
	
	
	http.HandleFunc("/hello",anyHandleFunc)

	http.ListenAndServe("0.0.0.0:8000",nil);
}