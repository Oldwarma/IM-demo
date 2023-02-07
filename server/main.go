package main

import (
	config2 "IM/server/config"
	"IM/server/internal1"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "E:\\GoSrc\\IM\\etc\\im.yaml", "the config2 file")

func main() {
	var c config2.Config
	conf.MustLoad(*configFile, &c)
	server := internal1.NewServer(c.TCPServer.Ip, c.TCPServer.Port)
	server.Start()

}
