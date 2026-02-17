package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Mini message queue : ")
	myBroker := &Broker{
		Subscribers: make(map[string][]*Subscriber),
	}
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn, myBroker)
	}
}

func handleConnection(conn net.Conn, b *Broker) {
	defer conn.Close()
	sub := &Subscriber{Conn: conn}
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Bytes()
		var msg Message
		err := json.Unmarshal(line, &msg)
		if err != nil {
			fmt.Printf("Invalid JSON received!\n")
			continue
		}
		fmt.Printf("Parsed successfully -> Command: %s | Topic: %s | Payload: %s\n", msg.Command, msg.Topic, msg.Payload)

		switch msg.Command {
		case "SUB":
			b.Subscribe(msg.Topic, sub)
		case "PUB":
			b.Publish(msg.Topic, msg.Payload)
		}
	}

	fmt.Printf("client disconnected!\n")
}
