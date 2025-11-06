package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	msg string
}

func (m model) Init()tea.Cmd{
	return nil
}

func(m model) Update(msg tea.Msg)(tea.Model,tea.Cmd){
	switch msg:=msg.(type){

	// kya key press hui agar hui to , key ko map kro
	case tea.KeyMsg:
		
		switch msg.String(){

		case "ctrl+c":
			fmt.Println("user clicked : ",msg.String());
			return m,tea.Quit;
		}

	}
	return m,nil;


}

func (m model) View() string{
	// color ansi or hex
	style:=lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("16")).Background(lipgloss.Color("205"))

	welcome:=style.Render("Welcome to Easy Trading")

	return welcome;
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
}