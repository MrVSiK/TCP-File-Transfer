package main

import (
	ftp "github.com/jlaffaye/ftp"
	"log"
	"time"
)

func main(){
	c, err := ftp.Dial("localhost:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("anonymous", "anonymous")
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the FTP conn

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}