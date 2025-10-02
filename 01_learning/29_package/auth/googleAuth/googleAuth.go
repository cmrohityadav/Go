package anyPackageName

import (
	"fmt"

	"github.com/cmrohityadav/go/01_learning/29_package/auth"
	"github.com/cmrohityadav/go/01_learning/29_package/session"
)

func GoogleLogin(email,password string){
	fmt.Println("Using google Auth but underhood we are using normal login HEHEhe......")

	fmt.Println("getting session")
	sSession:=session.Session();
	fmt.Println("Gotted session id: ",sSession);

	auth.LoginWithCredentials(email,password);
	
	
	fmt.Println("SuccessFully");
}
