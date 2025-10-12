package updatemenu

import (
	"fmt"

	"github.com/cmrohityadav/go/02_protogo/03_todocli/internal/cli/deletemenu"
)


func UpdateMenu(){
	fmt.Println("Update by ID");
	fmt.Println("1. Update Status");
	fmt.Println("2. Update Title");
	fmt.Println("3. Delete By ID");
	fmt.Println("0. Back");
	fmt.Print("Enter Valid Option: ");
	var op int;
	fmt.Scan(&op);

	switch op{
	case 1:
		UpdateStatus();
	case 2:
		UpdateTitle();

	case 3:
		deletemenu.DeleteMenu();
	case 0:
		return;
		
	default:
		fmt.Println("Enter Valid Option");
		UpdateMenu();

}



}