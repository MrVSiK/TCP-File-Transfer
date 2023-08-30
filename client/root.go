package main

import (
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "9001"
	TYPE = "tcp"
)

func main(){
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST + ":" + PORT)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte("Hello World"))

	if err != nil {
		log.Fatal(err)
	}

	received := make([]byte, 1024)

	_, err = conn.Read(received)

	if err != nil {
		log.Fatal(err)
	}

	println(string(received))
}