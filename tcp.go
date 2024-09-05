package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

const (
	Method = "tcp"
	Port = ":8080"
)

var (
	messages = []string{}
	Conns = map[string]net.Conn{}
	Addrs = map[string]string{}
)

func main() {
	listner, err := net.Listen(Method, Port)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		Name := make([]byte, 1024)
		con, err := listner.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		con.Write([]byte("write user name : "))
		i, err := con.Read(Name)
		if err != nil {
			continue
		}
		Name = Name[:i-1]
		msg := fmt.Sprintf("%s enter the chat", Name)
		printMessage([]byte(msg), string(Name))
		Conns[string(Name)] = con
		writeOldMessages(string(Name))
		go con_handler(string(Name))
	}
}

func writeOldMessages(name string) {
	for _, message := range messages {
		print_Message_to_con([]byte(message), name)
	}
}

func con_handler(name string) {
	con := Conns[name]
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		message := scanner.Text()
		msg := []byte(fmt.Sprintf("[%s]: %s", name, message))
		messages = append(messages, string(msg))
		go printMessage(msg, name)
	}
	delete(Conns, name)
	printMessage([]byte(fmt.Sprintf("%s left the chat\n", name)), name)
}

func printMessage(msg []byte, name string) {
	msg = append(msg, '\n')
	year, month, day := time.Now().Date()
	hour, min, sec:= time.Now().Clock()
	message := fmt.Sprintf("[%d-%s-%d %d:%d:%d]%s", year, month, day, hour, min, sec, msg)
	for Name, con := range Conns {
		if Name == name {continue}
		con.Write([]byte(message))
	}
}

func print_Message_to_con(msg []byte, name string) {
	msg = append(msg, '\n')
	year, month, day := time.Now().Date()
	hour, min, sec:= time.Now().Clock()
	message := fmt.Sprintf("[%d-%s-%d %d:%d:%d]%s", year, month, day, hour, min, sec, msg)
	con := Conns[name]
	con.Write([]byte(message))
		
}
