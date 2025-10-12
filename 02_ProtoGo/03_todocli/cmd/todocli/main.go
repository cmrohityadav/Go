package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/cmrohityadav/go/02_protogo/03_todocli/internal/cli"
	"github.com/cmrohityadav/go/02_protogo/03_todocli/internal/storage"
)

func main(){
	fmt.Println("Welcome to Smart Cli Todo");
	todoLocalFile:=filepath.Join("storage","store.json");
	todos,err:=storage.Load(todoLocalFile);
	if err!=nil{
		fmt.Println("No existing todo file found, starting fresh:", err);
	}

	if todos!=nil {
		storage.GlobalTodo=todos;
	}



	c:=make(chan os.Signal,1);

	signal.Notify(c,os.Interrupt,syscall.SIGTERM);
	

	go func(){
		<-c
		fmt.Println("\n⚠️  Detected Ctrl + C, saving todos before exit...")
		if err := storage.Save(todoLocalFile); err != nil {
			fmt.Println("Failed to save todos:", err)
		} else {
			fmt.Println("✅ Todos saved successfully!")
		}
		os.Exit(0)
	}();

	cli.MainCLI();


	if err := storage.Save(todoLocalFile); err != nil {
		fmt.Println("Failed to save todos:", err);
	}

}