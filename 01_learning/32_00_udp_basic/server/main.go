package main

import (
	"fmt"
	"log"
	"net"
)

func main(){
	pNetUDPAddr,err:=net.ResolveUDPAddr("udp","localhost:8000");
	if err!=nil{
		log.Fatal("Error while  binding IP and Port");
	}

	pNetUDPConn,err:=net.ListenUDP("udp",pNetUDPAddr);
	if err!=nil{
		log.Fatal("Error while listening UDP Server");
	}
	
	fmt.Println("✅ UDP Server listening on", pNetUDPAddr);

	defer pNetUDPConn.Close();

	buffer:=make([]byte,1024);

	for{
		
		//  --BLOCKS-- here until data arrives
		numberOfByteRead,pNetUDPAddrClient,err:=pNetUDPConn.ReadFromUDP(buffer);
		if err!=nil{
			log.Print("Error while reading data from client: ",err);
			continue;
		}

		fmt.Println("Recieved From :",pNetUDPAddrClient.IP);
		fmt.Println("Recieved Data :",string(buffer[:numberOfByteRead]));

 		//  usually fast, but can  --BLOCKS-- if OS buffer is full
		numberOfByteWrite,err:=pNetUDPConn.WriteToUDP([]byte("This is msg From UPD Server"),pNetUDPAddrClient);
		if err!=nil{
			log.Print("Fail to send msg to client :",pNetUDPAddrClient.IP,"Error : ",err);
		}

		log.Print("Number of Byte sended: ",numberOfByteWrite);

	}

}