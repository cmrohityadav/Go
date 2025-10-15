package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	
	netConnServer, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln("Failed to connect to server:", err)
	}

	defer netConnServer.Close()

	fmt.Println("Connected to server at localhost:8000")
	fmt.Println("Type messages and press Enter to send.")

	
	go func() {
		reader := bufio.NewReader(netConnServer)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Server connection closed:", err)
				os.Exit(0)
			}
			fmt.Print("Received: " + message)
		}
	}()

	// Main loop: read from stdin and send to server
	stdReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You: ")
		text, err := stdReader.ReadString('\n')
		if err != nil {
			log.Println("Error reading input:", err)
			continue
		}

		_, err = netConnServer.Write([]byte(text))
		if err != nil {
			log.Println("Error sending message:", err)
		}
	}
}
