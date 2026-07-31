package main

import (
	"auth/internal/httpserver"
	"log"
)

func main() {

	var ptrServerObject *httpserver.MyHTTPServerApp;
	ptrServerObject=&httpserver.MyHTTPServerApp{}
	err:=ptrServerObject.NewSetupServer();
	if err!=nil{
		log.Fatalf("Unable Setup Server: %v", err)
	}

	defer ptrServerObject.DB.Close()

	err=ptrServerObject.Start()
	if err!=nil{
		log.Fatal(err)
	}


}