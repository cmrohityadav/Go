package main

import "fmt"

func main() {
	fmt.Println("Welcome to Struct")
	/*
		No inheritance
		No super
		No parent
	*/

	rohit := User{"Rohit", "rohit@go.in", true, 24}

	fmt.Println(rohit)
	fmt.Printf("Rohit's Details are: %+v \n", rohit)
	fmt.Printf("Name is %v and email is %v", rohit.SName, rohit.SEmail)


    rohit.GetStatus()

    rohit.NewMail()
    fmt.Println(rohit)
    // not chnaging value due to its pass by value

}

type User struct {
	SName   string
	SEmail  string
	BStatus bool
	IAge    int

}

func (u User) GetStatus(){
    fmt.Println("\nIs user active : ",u.BStatus)
}

func (u User)NewMail(){
      u.SEmail="new@mail.go"  
      fmt.Println("new email is ",u.SEmail)
}    
