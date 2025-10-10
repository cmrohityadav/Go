package main

import (
	"fmt"
	"path/filepath"

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


	cli.MainCLI();


	if err := storage.Save(todoLocalFile); err != nil {
		fmt.Println("Failed to save todos:", err);
	}

}