package main

import (
	"crypto/tls"
	"log"
	"net"
	"time"
)

func main(){
	cer, err := tls.LoadX509KeyPair("server.crt", "server.key");
	if err != nil {
		log.Fatal(err);
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}};
	ln, err := tls.Listen("tcp", "localhost:9090", config);

	if err != nil {
		log.Fatal(err);
	}

	defer ln.Close();

	for {
		conn, err := ln.Accept();
		if err != nil {
			log.Fatal(err);
		}
		go handleConnection(conn);
	}
}

func handleConnection(conn net.Conn){
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