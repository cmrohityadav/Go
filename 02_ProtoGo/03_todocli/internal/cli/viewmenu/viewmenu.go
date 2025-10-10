package viewmenu

import (
	"fmt"

	"github.com/cmrohityadav/go/02_protogo/03_todocli/internal/storage"
)


func ViewMenu(){
	fmt.Println("\n==== View Menu ====");
	storage.GlobalTodo.View();
}