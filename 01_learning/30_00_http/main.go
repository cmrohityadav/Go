package main

import (
	"fmt"
	"net/http"
)

type anyHandlerObject struct{}

func (a anyHandlerObject)ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("This from object handler"))
}

func anyHandler() http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("rohit","yadav")
		w.Write([]byte("Rohit hanlder"))
	}) 
}

func anyHandlerFunc(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("using handlerFunc"))
}

func anyHandleFunc(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("using HandleFunc"))
}

func main() {
	fmt.Println("Hello");

	http.Handle("/anyhandler",anyHandler());
	http.Handle("/anyhandlerObject",anyHandlerObject{});
	http.Handle("/anyhandlerfunc",http.HandlerFunc(anyHandlerFunc)); 

	http.HandleFunc("/anyhandlefunc",anyHandleFunc)

	http.ListenAndServe("0.0.0.0:8080",nil);

}