package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/02_protogo/02_urlshortner/internal/config"
	"github.com/02_protogo/02_urlshortner/internal/http/handler"
)



func main(){
	fmt.Println("Welcome to url shortner")

	// ------------LOAD CONFIG---------
	cfg:=config.LoadConfig();

	//---------------IN- MEMORYDB---------
	
	
	//--------SETUP SERVER------------
	router:=http.NewServeMux();

	router.Handle("GET /",http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello this is the home page"));
	}))

	router.Handle("POST /create",handler.CreateShortUrl(cfg));


	router.HandleFunc("GET /{id}",handler.RedirectHits);

	server:=http.Server{
		Addr:fmt.Sprintf("%s:%d", cfg.ServerIp, cfg.Port) ,
		Handler: router,
	}

	log.Printf("Starting server on %s:%d...\n", cfg.ServerIp, cfg.Port)

	if err:=server.ListenAndServe(); err!=nil{
		log.Fatalf("Server failed: %v", err)
	}


}