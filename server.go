package main

import (
	"fmt"
	"net"
)

const (
	network = "tcp"
)

type Socket interface {
	Start()
}

type Server struct {
	Host, Port string
}

func NewServer(host, port string) Server {
	return Server{
		Host: host,
		Port: port,
	}
}

func (instance Server) Start() {
	address := fmt.Sprintf("%s:%s", instance.Host, instance.Port)

	listener, err := net.Listen(network, address)
	if err != nil {
		fmt.Println(err)
	}

	for true {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print(err)
		}

		fmt.Println("Connected by ", listener.Addr())

		data := make([]byte, 0, 4096)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Print(err)
		}

		fmt.Println("Data: ", data, n)

		_, err = conn.Write(data)
		if err != nil {
			fmt.Print(err)
		}

		conn.Close()
	}
}
