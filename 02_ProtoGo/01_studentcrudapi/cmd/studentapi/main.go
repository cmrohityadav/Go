package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"main/internal/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

done:= make(chan os.Signal,1)
signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM);
go func(){

	err:=server.ListenAndServe();

	fmt.Println("Server started");

	if err!=nil{
		log.Fatal("failed to start server");
	}

}()

<-done

slog.Info("Shutting down the server");

ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)

defer cancel();

err:=server.Shutdown(ctx);
if err!=nil{
	slog.Error("Failed to shutdown server",slog.String("error",err.Error()));
}


slog.Info("server shutdown succesfully");

}