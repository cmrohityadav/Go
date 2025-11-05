package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	msg string
}

func (m model) Init()tea.Cmd{
	return nil
}

func(m model) Update(msg tea.Msg)(tea.Model,tea.Cmd){
	return m,nil;
}

func (m model) View() string{
	return m.msg;
}


func initializeMode()model{
return model{
	msg: "biagan",
}
}

func main() {
	p:=tea.NewProgram(initializeMode());

	if _,err:=p.Run();err!=nil{
		fmt.Printf("Alas, There's been an error: %v",err);
		os.Exit(1);
	}
	fmt.Println("Welcome Totion using Bubble Tea")
}