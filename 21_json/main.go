package main

import (
	"encoding/json"
	"fmt"
)
type Course struct{
	Name string `json:"coursename"`
	Price int
	Platform string
	Password string `json:"-"` //just
	Tags [] string

}

func main() {
	fmt.Println("Welcome to JSON")
	EncodJson()

}

func  EncodJson(){
	myCourse:=[]Course{
		{"ReactJs",500,"youtube.com","123456",[]string{"reactjs","javascript","web development"}},
		{"C++",1000,"youtube.com","123456",[]string{"system","DSA"}},
		{"JavaScript",1500,"youtube.com","123456",nil},
	}

	finalJson,err:=json.MarshalIndent(myCourse,"","\t")
	if err!=nil{
		panic(err)

	}
	fmt.Printf("%s\n",finalJson)

}