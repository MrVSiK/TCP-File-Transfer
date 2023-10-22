package main

import (
	"crypto/tls"
	"encoding/binary"
	"encoding/hex"
	"log"
	"net"
	"os"
	"time"
	"fmt"
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
		go handleIncomingRequests(conn);
	}
}

func handleIncomingRequests(conn net.Conn){
	println("Received a request: " + conn.RemoteAddr().String());
	headerBuffer := make([]byte, 1024);

	_, err := conn.Read(headerBuffer);
	if err != nil {
		log.Fatal(err);
	}

	var name string;
	var reps uint32;

	if(headerBuffer[0] == byte(1) && headerBuffer[1023] == byte(0)){
		reps = binary.BigEndian.Uint32(headerBuffer[1:5]);
		lengthOfName := binary.BigEndian.Uint32(headerBuffer[5:9]);
		name = string(headerBuffer[9:9+lengthOfName]);
	} else {
		log.Fatal("Invalid header");
	}

	conn.Write([]byte("Header Received"));

	dataBuffer := make([]byte, 1024);

	file, err := os.Create("./received/" + name);

	if err != nil {
		log.Fatal(err);
	}

	for i := 0; i<int(reps); i++ {
		_, err := conn.Read(dataBuffer);
		if err != nil {
			log.Fatal(err);
		}

		if(dataBuffer[0] == byte(0) && dataBuffer[1023] == byte(1)){
			segmentNumber := dataBuffer[1:5];
			fmt.Printf("Segment Number: %d\n", binary.BigEndian.Uint32(segmentNumber));
			length := binary.BigEndian.Uint32(dataBuffer[5:9]);
			fmt.Printf("File Data: %s\n", hex.EncodeToString(dataBuffer[9:9+length]));
			file.Write(dataBuffer[9:9+length]);
		} else {
			log.Fatal("Invalid Segment");
		}
	}

	time := time.Now().UTC().Format("Monday, 02-Jan-06 15:04:05 MST");

	conn.Write([]byte(time));

	file.Close();
	conn.Close();
}