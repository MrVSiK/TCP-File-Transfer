package main

import (
	"log"
	"net"
	"time"
)

const (
	HOST = "localhost"
	PORT = "9001"
	TYPE = "tcp"
)

func main(){
	listen, err := net.Listen(TYPE, HOST + ":" + PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	println("Server has started on PORT " + PORT)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleIncomingRequests(conn)
	}
}

func handleIncomingRequests(conn net.Conn){
	println("Received a request: " + conn.RemoteAddr().String())
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	time := time.Now().UTC().Format("Monday, 02-Jan-06 15:04:05 MST")

	conn.Write([]byte(time))

	conn.Close()
}