package addmenu

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cmrohityadav/go/02_protogo/03_todocli/internal/storage"
)

func AddMenu() {

	fmt.Println("\n==== ADD MENU ====")
	for {
		fmt.Println("1. Add New Todo")
		fmt.Println("2. Back")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter Title: ")
			reader:=bufio.NewReader(os.Stdin);
			text,_:=reader.ReadString('\n')
			storage.GlobalTodo.ADD(text);
		case 2:
			return

		}

	}

}
