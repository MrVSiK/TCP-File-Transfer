package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"log"
	"os"
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

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "localhost:9090", conf)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

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
