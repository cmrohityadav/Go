package main

import (
	"fmt"

	"github.com/cmrohityadav/go/01_learning/29_package/auth"
	anyPackageName "github.com/cmrohityadav/go/01_learning/29_package/auth/googleAuth"
)


func main(){
auth.LoginWithCredentials("cmrohityadav","123456");

fmt.Println("In main using google Auth");

anyPackageName.GoogleLogin("cmrohityadav@gmail.com","rohit123")
}