package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	netcat "nc/outils"
)

const Usage = "[USAGE]: ./TCPChat $port"

func main() {
	if len(os.Args) > 2 {
		fmt.Println(Usage)
		return
	}
	port := "8989"
	if len(os.Args) == 2 && validPort(os.Args[1]) {
		port = os.Args[1]
	} else if len(os.Args) == 2 {
		fmt.Println(Usage)
		return
	}
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	Server := netcat.CreateNewServer()
	for {
		Conn, err := l.Accept()
		if err != nil {
			continue
		}
		go Server.WelcomeToTheServer(Conn)
	}
}

func validPort(s string) bool {
	num, err := strconv.Atoi(s)
	if err != nil || num < 1024 || num > 65535 {
		return false
	}
	return true
}
