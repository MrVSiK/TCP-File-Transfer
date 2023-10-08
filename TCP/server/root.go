package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"os"
	"time"
)

const (
	HOST = "localhost"
	PORT = "9001"
	TYPE = "tcp"
)

func getDataFromFile(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var fileData bytes.Buffer;
	fileData.Write(content);

	lengthData := make([]byte, 4);

	binary.BigEndian.PutUint32(lengthData, uint32(fileData.Len()));


	return append(lengthData, content...);

}

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
	println("Received a request: " + conn.RemoteAddr().String());
	buffer := make([]byte, 1024);

	_, err := conn.Read(buffer);
	if err != nil {
		log.Fatal(err);
	}

	lengthOfFileData := int(binary.BigEndian.Uint32(buffer[0:4]));

	fileData := buffer[4:4+lengthOfFileData];

	err = os.WriteFile("received.txt", fileData, 0644);
	if err != nil {
		log.Fatal(err);
	}

	time := time.Now().UTC().Format("Monday, 02-Jan-06 15:04:05 MST");

	conn.Write([]byte(time));

	conn.Close();
}