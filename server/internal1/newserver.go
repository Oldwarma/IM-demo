package internal1

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

type Server struct {
	Ip   string
	Port string

	Map map[string]*User
	Mu  sync.RWMutex

	//消息广播
	Msg chan string
}

func NewServer(ip, port string) *Server {
	return &Server{
		Ip:   ip,
		Port: port,
		Map:  map[string]*User{},
		Msg:  make(chan string),
	}

}

func (ns *Server) Start() {
	l, err := net.Listen("tcp", ns.Ip+":"+ns.Port)
	if err != nil {
		log.Println("listen err")
	}
	defer l.Close()
	for {
		conn, err1 := l.Accept()
		if err1 != nil {
			log.Println("Accept err")
		}
		fmt.Println(12313)

		ns.handler(conn)
	}
}

func (ns *Server) handler(conn net.Conn) {
	user := NewClient(conn)
	go ns.ListerMsg()
	//发送用户上线信息
	ns.OnLine(user)

	//活跃用户
	//isLive := make(chan bool)

	go func() {
		buf := make([]byte, 4099)
		for {
			read, err := conn.Read(buf[:])
			if read == 0 {
				ns.OffLine(user)
				return
			}
			fmt.Println("接收client:", string(buf[:read]))
			if err != nil && err != io.EOF {
				fmt.Println("read err")
			}
			msg := string(buf[:read])

			ns.DoMsg(user, msg)

			//isLive <- true
		}
	}()

	//for {
	//	select {
	//	case <-isLive:
	//	case <-time.After(time.Second * 10):
	//		user.SendMsg("你要被踢了")
	//
	//		return
	//	}
	//}
}

func (ns *Server) BroadCast(user *User, msg string) {
	sendMsg := "用户地址:" + user.Addr + "\n" + "用户名称" + user.Name + "\n" + "用户信息" + msg
	ns.Msg <- sendMsg
}

func (ns *Server) ListerMsg() {
	for {
		msg := <-ns.Msg
		ns.Mu.Lock()
		for _, user := range ns.Map {
			user.C <- msg
		}
		ns.Mu.Unlock()
	}
}
func (ns *Server) OnLine(user *User) {
	ns.Mu.Lock()
	ns.Map[user.Name] = user
	ns.Mu.Unlock()
	ns.BroadCast(user, " 已上线")
}
func (ns *Server) OffLine(user *User) {
	ns.Mu.Lock()
	ns.Map[user.Name] = user
	ns.Mu.Unlock()
	ns.BroadCast(user, " 已下线")
}

func (ns *Server) DoMsg(user *User, msg string) {
	if msg == "who" {
		//查询当前在线用户都有哪些
		ns.Mu.Lock()
		for _, u := range ns.Map {
			onlineMsg := "用户地址:" + u.Addr + "\n" + "用户名称" + u.Name + "\n" + "在线...\n"
			user.SendMsg(onlineMsg)
		}
		ns.Mu.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//消息格式:rename|张三
		newName := strings.Split(msg, "\r")[0]
		//判断name是否存在
		if ns.LookUser(newName) {
			user.SendMsg("当前用户名被使用\n")
		} else {
			ns.AddUser(user, newName)
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			user.SendMsg("消息格式不正确，请使用“to|张三|你好")
			return
		}
		remoteUser, ok := ns.Map[remoteName]
		if !ok {
			user.SendMsg("该用户名不存在\n")
			return
		}
		content := strings.Split(msg, "|")[2]
		if content == "" {
			user.SendMsg("无消息内容,请重发\n")
			return
		}
		remoteUser.SendMsg(user.Name + "对你说:" + content)
	} else {
		ns.BroadCast(user, msg)
	}
}

func (ns *Server) LookUser(name string) bool {
	_, ok := ns.Map[name]
	if ok {
		return true
	}
	return false
}
func (ns *Server) AddUser(user *User, newName string) {
	ns.Mu.Lock()
	defer ns.Mu.Unlock()
	delete(ns.Map, user.Name)
	user.Name = newName
	ns.Map[newName] = user
}
