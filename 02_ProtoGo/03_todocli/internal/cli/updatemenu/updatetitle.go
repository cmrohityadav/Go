package updatemenu

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cmrohityadav/go/02_protogo/03_todocli/internal/storage"
)

func UpdateTitle(){
	var id int;
	fmt.Print("Enter Id: ");
	fmt.Scan(&id);
	if valid:=storage.GlobalTodo.ValidateId(id); !valid{
		UpdateTitle();
	}
	
	fmt.Print("Enter new title : ");
	reader:=bufio.NewReader(os.Stdin);
	newTitle,err:=reader.ReadString('\n');
	if err!=nil{
		fmt.Println("ERROR READING CLI TEXT");
	}
	err=storage.GlobalTodo.UpdateTitleById(id,newTitle);
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("✅ Task updated successfully!")
	}

}

