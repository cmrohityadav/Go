package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"main/internal/config"
	"main/internal/http/handlers/student"
	"main/internal/storage/sqlite"
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
storage,err1:=sqlite.New(cfg);
if err1!=nil {
	log.Fatal(err1);
}

slog.Info("Storage initialized",slog.String("Env",cfg.Env));


// ----------setup router-----------
router:=http.NewServeMux();

router.HandleFunc("GET /",func (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welocome to HOME students api"));
})

router.HandleFunc("POST /api/student",student.New(storage));




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