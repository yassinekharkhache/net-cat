package netcat

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

//this function write the old messages in the chat to the new user
func writeOldMessages(name string) {
	for _, message := range Server.Messages {
		print_Message_to_con([]byte(message), name)
	}
}

//this function take two variables the name of the user and the textmessage 
//and format it to be in the right message formula
func formatMessages(name, text string) []byte {
	msg := fmt.Sprintf("[%s]: %s\n", name, text)
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	return []byte(fmt.Sprintf("[%d-%s-%d %d:%d:%d]%s", year, month, day, hour, min, sec, msg))
}

//that function handle the connextion for a single user by his name
func con_handler(name string) {
	Conn := Server.Conns[name]
	scanner := bufio.NewScanner(Conn)
	for scanner.Scan() {
		msg := formatMessages(name, scanner.Text())
		Server.Messages = append(Server.Messages, msg)
		go printMessage([]byte(msg), name)
	}
	deleteConnection(Conn)
	delete(Server.Conns, name)
	printMessage([]byte(fmt.Sprintf("%s left the chat\n", name)), name)
}

//delete the connection from the array
func deleteConnection(Conn net.Conn) {
	for i, con := range Server.Users {
		if con == Conn && i != len(Server.Users) {
			Server.Users = append(Server.Users[:i], Server.Users[i+1:]...)
			return
		} else if con == Conn {
			Server.Users = Server.Users[:len(Server.Users)-1]
		}
	}
}

// print the the message to the other users
func printMessage(msg []byte, name string) {
	for Name, con := range Server.Conns {
		if Name == name {
			continue
		}
		con.Write(msg)
	}
}

//print message to a specific user
func print_Message_to_con(msg []byte, name string) {
	con := Server.Conns[name]
	con.Write(msg)
}
