package main

import (
	"bufio"
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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		fmt.Printf("received: %s\n", scanner.Text())
	}

	fmt.Printf("client disconnected!\n")
}
