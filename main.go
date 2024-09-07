package main

import (
	"fmt"
	"net"
	"os"

	netcat "netcat/utils"
)

func main() {
	// Create new Socket
	listner, err := net.Listen(netcat.Method, netcat.Port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listner.Close()
	// Set the hello message
	netcat.Server.HelloMessage, err = os.ReadFile("linuxHello.txt")
	if err != nil {
		fmt.Println("error laoding the file")
		return
	}
	
	netcat.AcceptConnections(listner)
}
