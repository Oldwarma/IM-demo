package main

import (
	"IM/client/config2"
	"IM/client/internal2"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("c", "E:\\GoSrc\\IM\\etc\\im.yaml", "config2 file")

func main() {
	var c config2.Config
	conf.MustLoad(*configFile, &c)

	internal2.StartClient(c.TCPServer.Ip, c.TCPServer.Port)

}
