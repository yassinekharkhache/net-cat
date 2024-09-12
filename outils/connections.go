package nc

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

func (s *Server) GetUserInfo(Conn net.Conn) ([]byte, []byte, error) {
	Name := make([]byte, 50)
	group := make([]byte, 50)
	res, err := Conn.Read(Name)
	if err != nil || res == 1 || s.invalidName(Name[:res-1]) || res == 20 {
		Conn.Write([]byte("invalid name connection is closed\n"))
		if err == nil {
			err = errors.New("invalid name")
		}
		return nil, nil, err
	}
	Name = Name[:res-1]
	Conn.Write([]byte("Choose your group chat:\n"))
	for group := range s.Groups {
		Msg := fmt.Sprintf(". %s\n", group)
		Conn.Write([]byte(Msg))
	}
	Conn.Write([]byte("Name Of Groupe : "))
	res, err = Conn.Read(group)
	if err != nil || res == 1 || s.invalidGrpoup(group[:res-1]) || res == 50 {
		Conn.Write([]byte("invalid group connection is closed\n"))
		if err == nil {
			err = errors.New("invalid group")
		}
		return nil, nil, err
	}
	group = group[:res-1]
	if len(s.Groups[string(group)]) > 2 {
		Conn.Write([]byte("Groupe already have 10 people"))
		return nil, nil, err
	}
	return Name, group, nil
}

func (s *Server) invalidName(name []byte) bool {
	if InvalidMsg(name) {
		return true
	}
	for _, Group := range s.Groups {
		for Name := range Group {
			if Name == string(name) {
				return true
			}
		}
	}
	return false
}

func InvalidMsg(name []byte) bool {
	if len(strings.TrimSpace(string(name))) > 0 {
	}
	for _, v := range name {
		if v > 126 || v < 32 {
			return true
		}
	}
	return false
}

func (s *Server) WriteOldMessages(Conn net.Conn, group string) {
	for _, message := range s.Messages[group] {
		Conn.Write([]byte(message))
	}
}

func (s *Server) WriteMessage(Name string, group string, message string) {
	s.mutex.Lock()
	s.Messages[group] = append(s.Messages[group], message)
	s.mutex.Unlock()
	for UserName, UserConn := range s.Groups[group] {
		if UserName != Name {
			UserConn.Write([]byte("\n" + message))
			UserConn.Write([]byte(Format("", string(UserName))))
		}
	}
}

func Format(msg, name string) string {
	time := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s][%s]: %s", time, name, msg)
}

func (s *Server) invalidGrpoup(Grp []byte) bool {
	for group := range s.Groups {
		if string(Grp) == group {
			return false
		}
	}
	return true
}
