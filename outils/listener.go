package nc

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

type Users map[string]net.Conn

type Server struct {
	Groups   map[string]Users
	mutex    sync.Mutex
	Messages map[string][]string
	HelloMsg []byte
}

func CreateNewServer() *Server {
	groups := make(map[string]Users, 3)
	groups["General"] = Users{}
	groups["Random"] = Users{}
	groups["vip"] = Users{}
	hellMsg, err := os.ReadFile("linuxHello.txt")
	if err != nil {
		panic("error reading the file\n")
	}
	return &Server{
		Groups:   groups,
		Messages: make(map[string][]string, 3),
		HelloMsg: hellMsg,
	}
}

func (s *Server) WelcomeToTheServer(Conn net.Conn) {
	defer Conn.Close()
	Conn.Write(s.HelloMsg)
	Name, group, err := s.GetUserInfo(Conn)
	
	if err != nil {
		return
	}
	s.WriteOldMessages(Conn, string(group))
	s.mutex.Lock()
	// add new User
	s.Groups[string(group)][string(Name)] = Conn
	s.mutex.Unlock()
	// join Message
	JoinMessage := fmt.Sprintf("%s Joined The groupe chat\n", string(Name))
	s.WriteMessage(string(Name), string(group), JoinMessage)
	s.StartChat(string(Name), string(group))
	DisconnectMsg := fmt.Sprintf("%s Leave The groupe chat\n", string(Name))
	s.WriteMessage(string(Name), string(group), DisconnectMsg)
	s.mutex.Lock()
	delete(s.Groups[string(group)], string(Name))
	s.mutex.Unlock()
}

func (s *Server) StartChat(Name, group string) {
	Conn := s.Groups[group][Name]
	for {
		Conn.Write([]byte(Format("", Name)))
		msg := make([]byte, 200)
		index, err := Conn.Read(msg)
		if err != nil || index == 1 || InvalidMsg(msg[:index-1]) || index == 200 {
			Conn.Write([]byte("invalid message\n Connection denied\n"))
			return
		}
		if strings.HasPrefix(string(msg), "--name ") && s.invalidName(msg[7:]) {
			EditMessage := fmt.Sprintf("%s now is %s\n", Name, string(msg[7:]))
			s.mutex.Lock()
			delete(s.Groups[group], Name)
			// remove the flag and the newline
			Name = string(msg[7 : index-1])
			s.Groups[string(group)][Name] = Conn
			s.mutex.Unlock()
			s.WriteMessage(Name, group, EditMessage)
			continue
		}
		message := Format(string(msg), Name)
		s.WriteMessage(Name, group, message)
	}
}
