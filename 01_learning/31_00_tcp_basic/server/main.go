package main

import (
	"fmt"
	"log"
	"net"
)
func handleConnection(clientConnection net.Conn){
	defer clientConnection.Close();

	buffer:=make([]byte,1024);
	
	// 🟥 BLOCKING
	numberOfByteRead,err:=clientConnection.Read(buffer);
	if err != nil {
			fmt.Println("Client disconnected:", clientConnection.RemoteAddr())
            fmt.Println("Client disconnected:", err)
            return
    }

	msg := string(buffer[:numberOfByteRead]);
	fmt.Println("Received From Client:", msg);

	clientConnection.Write([]byte("Echo: " + msg));
}
func main(){
	// 🟢 NOT blocking
	netListner,err:=net.Listen("tcp","localhost:8000");
	if err!=nil{
		log.Fatalln(err);
	}

	defer netListner.Close();

	for{
		// accept an incoming connection
		// 🟥 BLOCKING
		netConnectionClient,err:=netListner.Accept();
		if err!=nil{
			log.Println(err);
		}

		// handle the connection
		go handleConnection(netConnectionClient);


	}
}