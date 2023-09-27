package main

import (
	"crypto/tls"
	"log"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "localhost:9090", conf)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

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
