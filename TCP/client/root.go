package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"os"
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
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST + ":" + PORT)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write(getDataFromFile("text.txt"))

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