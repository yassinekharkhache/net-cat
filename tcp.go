package main

import (
	"bufio"
	"fmt"
	"net"
)

var users []net.Conn

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	for {
		con, err := listner.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go con_handler(con)
	}
}

func con_handler(con net.Conn) {
	defer con.Close()
	con.Write([]byte("write user name : "))
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
