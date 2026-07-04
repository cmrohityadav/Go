package main

import (
	"fmt"
)

type User struct {
	Name string
	aAmt int
	status string
}

	//methods //value reciever
	func (objUser User) ChangeUserStatus(){
		objUser.status="Online";
		fmt.Println("In Method : ChangeUserStatus")
		fmt.Println(objUser)
	}
	func (objUser *User) pChangeUserStatus(){
		objUser.status="Online";
		fmt.Println("In Method : pChangeUserStatus")
		fmt.Println(objUser)
	}
func main() {
	var Rohit User;
	Rohit.aAmt=10
	Rohit.Name="rohit"
	fmt.Println(Rohit)

	rahul:=&User{
		Name: "rohit",
		aAmt: 50,
	}

	fmt.Println(rahul)

	piyush:= User{aAmt: 40,Name: "piyush"}
	fmt.Println(piyush)

	vivek:=User{"vivek",60,"offline"}
	fmt.Println(vivek)

	var sonam User=User{"sonam",70,"offline"}
	fmt.Println(sonam)

	sonam.ChangeUserStatus()
	fmt.Println(sonam)
	
	sonam.pChangeUserStatus()
	fmt.Println(sonam)


	
	

}