package cli

import (
	"fmt"
	"github.com/cmrohityadav/go/02_protogo/03_todocli/internal/cli/addmenu"
	"github.com/cmrohityadav/go/02_protogo/03_todocli/internal/cli/viewmenu"
	"github.com/cmrohityadav/go/02_protogo/03_todocli/internal/constant"
)


func MainCLI(){

		
	for {
		fmt.Println("\n==== Main Menu ====");
		fmt.Println("1. View");
		fmt.Println("2. ADD");
		fmt.Println("3. UPDATE");
		fmt.Println("4. DELETE");
		fmt.Println("5. EXIT");

		var selected int;
		fmt.Print("\nEnter: ");
		fmt.Scan(&selected);

		switch constant.OptionsMainMenu(selected){
		case constant.VIEW:
			viewmenu.ViewMenu();
		case constant.ADD:
			addmenu.AddMenu();
		case constant.DELETE:
			fmt.Println("DELETE");
		case constant.UPDATE:
			fmt.Println("UPDATE");
		case constant.EXIT:
			fmt.Println("EXIT....");
			return
		default:
			fmt.Println("Invalid option");

		}
	}
}