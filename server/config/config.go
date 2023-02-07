package config

type Config struct {
	TCPServer TCPServer
}
type TCPServer struct {
	Ip   string
	Port string
}
