package types

import (
	"net"
	"sync"
)

type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn
}

type Server struct {
	Ip   string
	Port string

	Map map[string]*User
	mu  sync.RWMutex

	//消息广播
	msg chan string
}
