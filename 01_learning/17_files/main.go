package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to files in golang")

	content:="This needs to go in file - cmrohityadav.in"

	anyFile,error:=os.Create("./fileCreatedByGo.txt")

	if error!=nil{
		panic(error)
	}

	length,err:=io.WriteString(anyFile,content)
	
	if err !=nil{
		panic(err)
	}
	fmt.Println("lenght is : ",length)

	defer anyFile.Close()


	readFile("./fileCreatedByGo.txt")

}

func readFile(filename string){
	itsDataInByte,err:=ioutil.ReadFile(filename)
	if err !=nil{
		panic(err)
	}

	fmt.Println("Text data inside the file (Raw) is \n",itsDataInByte)
	fmt.Println("Text data inside the file (string) is \n",string(itsDataInByte))
}