package main

<<<<<<< HEAD
import "fmt"

func main() {

	fmt.Println("rohit yadav ji")

	var fullname string;

	fullname="Rohit C Yadav"
	fmt.Println("my full name address ",&fullname);
	fmt.Println("my full name value ",fullname);


	var mobile int;


	fmt.Println("my mobile number : ",mobile);
	fmt.Println("address of mobile no: ",&mobile);

	var city string="mumbai"

	fmt.Println("This is my city: ",city)
	fmt.Printf("This is my city %s",city);


	var amount float64=10.00;

	fmt.Println("\nThis is my amount: ",amount)
	fmt.Printf("this is my amout by format speficiers %.4f",amount);

	surname:=99.0001000;
	fmt.Println("surname :",surname);
	fmt.Printf("The datatype of name is %v",surname);



=======
import (
	"fmt"
	"unsafe"
)

func main() {
	var name uint8;
	fmt.Println(unsafe.Sizeof(name));
>>>>>>> 60613b61719b6b3b8ff2cd066b093e0575cde009
}