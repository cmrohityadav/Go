package main

import (
	"fmt"
	"net"
)

func main(){
	// BLOCKING 
	netConnection,err:=net.Dial("tcp","localhost:8000");
	if err!=nil{
		fmt.Println("Error connecting to server:", err)
        return
	}

	defer netConnection.Close();

	fmt.Println("Connected to server!");

	msg:="Hello I am Client";

	// BLOCKING 
	_,err=netConnection.Write([]byte(msg));
	if err!=nil{
		fmt.Println("Error sending:", err)
        return
	}

	buffer:=make([]byte,1024);

	// BLOCKING 
	numberOfByteRead,err:=netConnection.Read(buffer);

	if err!=nil{
		fmt.Println("Error reading:",err);
		return;
	}

 	fmt.Println("Server says:", string(buffer[:numberOfByteRead]));


}
