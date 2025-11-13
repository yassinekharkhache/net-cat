package main

import (
	"bufio"
	"fmt"
	"net"
)

var Users = make(map[string]net.Conn)

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	for {
		Name := []byte{}
		con, err := listner.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		con.Write([]byte("write user name : "))
		_, err = con.Read(Name)
		if err != nil {
			continue
		}
		Users[string(Name)] = con
		go con_handler(string(Name))
	}
}

func con_handler(name string) {
	con := Users[name]
	defer con.Close()
	scanner := bufio.NewScanner(con)
	msg := []byte{}
	for scanner.Scan() {
		con.Read(msg)
		go printMessage(string(msg), name)
	}
}

func printMessage(msg, name string) {
	for _, con := range Users {
		// message := []byte(fmt.Sprintf("[%s]: %s", name, msg))
		con.Write([]byte("kteb hadxi and kolxi"))
	}
}
