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
		UpdatedAt: time.Time{},
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

		fmt.Printf("ID: %d | Status: %s | CreatedAt: %v  | UpdatedAt: %v\n",myTodo.ID,myTodo.Status,myTodo.CreatedAt,myTodo.UpdatedAt);
		fmt.Println("Task: ",myTodo.Title);

	}
}

