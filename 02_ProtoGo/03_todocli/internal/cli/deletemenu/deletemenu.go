package deletemenu

import (
	"fmt"

	"github.com/cmrohityadav/go/02_protogo/03_todocli/internal/storage"
)


func DeleteMenu(){
	fmt.Println("1. are you sure to Delete");
	fmt.Println("0. Back");

	var op int;
	fmt.Print("Enter : ");
	fmt.Scan(&op);
	switch op{

	case 1:
		err:=DeleteById();
		if err!=nil{
			fmt.Println(err);
			DeleteMenu();
		}
		
	case 0:
		return;

	default:
		fmt.Println("Enter a valid opton");
		DeleteMenu();
	}


}

func DeleteById()error{
	var id int;
	fmt.Print("Enter ID: ");
	fmt.Scan(&id);

	err:=storage.GlobalTodo.DeleteById(id);
	if err!=nil {
		return err;
	};

	return nil;
}