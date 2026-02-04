package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Mini message queue : ")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("client connected")
		conn.Close()
	}
}
