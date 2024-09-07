package main

import (
	"fmt"
	"net"
	"os"

	netcat "netcat/utils"
)

func main() {
	// Create new Socket
	Port := "8989"
	if len(os.Args) > 2 {
		fmt.Println(netcat.Usage)
		return
	}
	if len(os.Args) == 2  && netcat.IsValidPort(os.Args[1]){
		Port = os.Args[1]
	}else if len(os.Args) == 2{
		fmt.Println(netcat.Usage)
		return
	}

	listner, err := net.Listen(netcat.Method, fmt.Sprintf(":%s", Port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listner.Close()
	Server := netcat.Startserver()
	// Set the hello message
	Server.HelloMessage, err = os.ReadFile("linuxHello.txt")
	if err != nil {
		fmt.Println("error laoding the file")
		return
	}
	Server.AcceptConnections(listner)
}
