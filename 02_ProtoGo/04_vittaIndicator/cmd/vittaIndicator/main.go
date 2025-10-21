package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"vittaindicator/internal/autodownload"
	"vittaindicator/internal/config"
)


func main() {

	fmt.Println("Welcome to Vitta Indicator");

	sWorkingDir,err:=os.Getwd();
	if err!=nil{
		log.Fatal("Get Fail to config Path Please provide config file path: ",err);
	}
	sConfigFilePath:=filepath.Join(sWorkingDir,"config","config.json");
	var cfg config.Config
	err=cfg.LoadConfig(sConfigFilePath);
	if err!=nil{
		log.Fatal("unable to load config file:",err)
	}

	go func(){
		autodownload.AutoDownload(cfg);
	}()

	quite:=make(chan os.Signal,1);

	signal.Notify(quite,os.Interrupt,syscall.SIGTERM,syscall.SIGINT);
	<-quite;

	
}