package internal2

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int //当前client的模式
}

func StartClient(ip, port string) {
	c, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		log.Println("conn err")
	}
	fmt.Println("client 已连接")
	go Reader(c)
	//Send(c)

	go Resp(c)
	for {

	}

	defer c.Close()
}

func Reader(c net.Conn) {
	for {
		//reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		var remoteName string
		fmt.Scanln(&remoteName)
		fmt.Println("发送:", remoteName)
		fmt.Fprintf(c, remoteName)
	}

}
func Resp(c net.Conn) {
	for {
		reader, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Println("接收服务端:" + reader)
	}

}
func Send(c net.Conn) {
	//resp, _ := bufio.NewReader(c).ReadString('\n')
	//fmt.Printf("接收:" + resp + "\n")
	sendMsg := "who\n"
	_, err := c.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err", err)
		return
	}
	fmt.Println("查询在线用户")
}
