package main
import "fmt"
func main(){
	var age int;

	fmt.Print("Enter your age: ");
	fmt.Scan(&age);

	if age<18{
		fmt.Println("your age>18: ",age);
		fmt.Println("your age>18: ",age);
		
	}else if age>25{
		fmt.Println("your age>25: ",age);
	}else if age>30{
		fmt.Println("your age>30: ",age);
	}else{
		fmt.Println("your age>: ",age);
	}
}