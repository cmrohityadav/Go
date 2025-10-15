package main

import (
	"fmt"
	"log"
	"net"
)

func main(){

	pNetUDPAddrServer,err:=net.ResolveUDPAddr("udp","localhost:8000");
	if err!=nil{
		log.Fatal("Error while binding port and Ip",err);
	}

	pNetUDPConn,err:=net.DialUDP("udp",nil,pNetUDPAddrServer);
	if err!=nil{
		log.Fatal("Error while Dialing to Server :",err);
	}

	defer pNetUDPConn.Close();

	numberOfByteWrite,err:=pNetUDPConn.Write([]byte("Hello This message from Client UDP "));
	if err!=nil{
		log.Print("Fail to Send message to Server UDP From Clinet UDP: ",err);
	}

	log.Println("Number of byte send to UPD Server :",numberOfByteWrite);

	buffer := make([]byte, 1024)
	numberOfByteRead, err := pNetUDPConn.Read(buffer) 
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}
	fmt.Println("Number of Byte Recieved :",numberOfByteRead);
	fmt.Printf("Received from server: %s\n", string(buffer[:numberOfByteRead]));



}