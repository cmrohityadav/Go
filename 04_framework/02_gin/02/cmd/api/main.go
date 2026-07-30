package main

import (
	"auth/internal/httpserver"
	"log"
	"net/http"
)

func main() {

	router := httpserver.NewRouter()

	ptrServer:=&http.Server{
		Addr: ":3000",
		Handler: router,
	}
	log.Printf("Serving Going to Start on Port : %s , Please wait...",ptrServer.Addr)
	if err:=ptrServer.ListenAndServe();err!=nil{
		if err==http.ErrServerClosed{
			log.Printf("Server Closed")
			return
		}
		log.Fatalf("Server error: %v",err)
	}
}