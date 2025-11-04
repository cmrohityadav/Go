package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

var Client = make(map[net.Conn]string)
var MyMutex sync.Mutex

func main() {
	fmt.Println("Go server")
	pNetListener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatalln("Server Fail to listen while binding with port and IP")
	}

	log.Println("Server Started at :", pNetListener.Addr().String())

	defer pNetListener.Close()

	for {

		pNetConnection, err := pNetListener.Accept()
		if err != nil {
			log.Println("Unable to establish connection with ", pNetConnection.RemoteAddr().String(), "Error :", err)
			continue
		}

		go handleClientNetConnection(pNetConnection)

	}

}

func handleClientNetConnection(clientNetConnection net.Conn) {
	defer clientNetConnection.Close()
	MyMutex.Lock()
	Client[clientNetConnection] = clientNetConnection.RemoteAddr().String()
	MyMutex.Unlock()

	log.Println("Client connected : ", clientNetConnection.RemoteAddr())

	buffer := make([]byte, 1024)

	for {
		numberOfByteRecv, err := clientNetConnection.Read(buffer)
		if err != nil {
			log.Println("Error while receiving Data: ", err)
			MyMutex.Lock()
			delete(Client, clientNetConnection)
			MyMutex.Unlock()
			break
		}

		BroadCastToEveryClient(buffer[:numberOfByteRecv], clientNetConnection)
	}

}

func BroadCastToEveryClient(buffer []byte, netConn net.Conn) {
	MyMutex.Lock();
	clients := make([]net.Conn, 0, len(Client))
	for c := range Client {
		if c != netConn {
			clients = append(clients, c)
		}
	}
	MyMutex.Unlock();

	for _, c := range clients {
		_, err := c.Write(buffer)
		if err != nil {
			log.Println("Error while broadcasting to ", c.RemoteAddr())
			c.Close()
			MyMutex.Lock()
			delete(Client, c)
			MyMutex.Unlock()
		}
	}

}
