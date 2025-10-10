package types

import (
	"fmt"
	"time"
)

type Todo struct{
	ID int
	Title string
	Status string
	CreatedAt time.Time
	UpdatedAt  time.Time
}

type Todos struct{
	List []Todo
}

func (t *Todos) ADD(title string){
	var temTodo=Todo{
		ID: len(t.List)+1,
		Title: title,
		Status: "pending",
		CreatedAt: time.Now(),
	}

	t.List=append(t.List, temTodo);
	fmt.Println("✅ Task added successfully!");
}

func (t *Todos) View(){
	if len(t.List) == 0 {
		fmt.Println("📭 No tasks available! Please add some.");
		return;
	}

	fmt.Println("\n==== Task List ====")
	

	for _, myTodo := range t.List {
		fmt.Println("-------------------------------------------------------------------------");
		createdAt:=myTodo.CreatedAt.Format("02-01-2006 15:04:05");

		updateAt:="-";
		if !myTodo.UpdatedAt.IsZero(){
			updateAt=myTodo.UpdatedAt.Format("02-01-2006 15:04:05");
		}


		fmt.Printf("ID: %d | Status: %s | CreatedAt: %s  | UpdatedAt: %s\n",myTodo.ID,myTodo.Status,createdAt,updateAt);
		fmt.Println("Task: ",myTodo.Title);

	}
	
}
func (t *Todos) ValidateId(id int) bool{
	if id <=len(t.List) && id>=1{
		return true;
	}
	return false;
}

func (t *Todos)UpdateTitleById(id int,newTitle string)error{
	
	if !t.ValidateId(id){
		return fmt.Errorf("please Enter Valid Id");
	}

	t.List[id-1].Title=newTitle;
	t.List[id-1].UpdatedAt=time.Now();

	return nil;
}


func (t *Todos) CompleteById(id int) error {
	if !t.ValidateId(id) {
		return fmt.Errorf("please enter a valid Id")
	}
	t.List[id-1].Status = "done"
	t.List[id-1].UpdatedAt = time.Now()
	return nil
}

