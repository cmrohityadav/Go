package main

import (
	"fmt"
	"log"
	"main/internal/config"
	"net/http"
)

func main(){
fmt.Println("Welcome student API");

// ----------load Config---------
cfg:=config.MustLoad()
	
// ----------Database setup---------


// ----------setup router-----------
router:=http.NewServeMux();

router.HandleFunc("GET /",func (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welocome to HOME students api"));
})




// ----------setup server---------
server:=http.Server{
	Addr: cfg.Addr,
	Handler: router,

}

err:=server.ListenAndServe();
fmt.Println("Server started");
if err!=nil{
	log.Fatal("failed to start server");
}


}