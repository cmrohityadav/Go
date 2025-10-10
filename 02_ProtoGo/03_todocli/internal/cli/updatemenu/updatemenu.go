package updatemenu

import (
	"fmt"

)


func UpdateMenu(){
	fmt.Println("Update by ID");
	fmt.Println("1. Update Status");
	fmt.Println("2. Update Title");
	fmt.Print("Enter Valid Option: ");
	var op int;
	fmt.Scan(&op);

	switch op{
	case 1:
		UpdateStatus();
	case 2:
		UpdateTitle();
	default:
		fmt.Println("Enter Valid Option");
		UpdateMenu();

}



}