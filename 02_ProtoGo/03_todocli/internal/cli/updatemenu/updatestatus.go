package updatemenu

import (
	"fmt"

	"github.com/cmrohityadav/go/02_protogo/03_todocli/internal/storage"
)
func UpdateStatus() {
	var id int;
	fmt.Print("Enter Id: ");
	fmt.Scan(&id);
	
	if valid:=storage.GlobalTodo.ValidateId(id); !valid{
		UpdateStatus();
	}
	
	err := storage.GlobalTodo.CompleteById(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("✅ Task marked as done!")
	}
}