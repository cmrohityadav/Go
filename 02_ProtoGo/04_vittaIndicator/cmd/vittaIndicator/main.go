package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"vittaindicator/internal/config"
)
var Config config.Config;

func main() {

	fmt.Println("Welcome to Vitta Indicator");

	sWorkingDir,err:=os.Getwd();
	if err!=nil{
		log.Fatal("Get Fail to config Path Please provide config file path: ",err);
	}
	sConfigFilePath:=filepath.Join(sWorkingDir,"config","config.json");

	err=Config.LoadConfig(sConfigFilePath);
	if err!=nil{
		log.Fatal("unable to load config file:",err)
	}
	

	
}