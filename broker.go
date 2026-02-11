package main

import (
	"fmt"
	"net"
)

type Message struct {
	Command string `json:"command"`
	Topic   string `json:"topic"`
	Payload string `json:"payload"`
}

type Broker struct {
	// Topic -> connections subscribed to that topic
	Subscribers map[string][]net.Conn
}

func (b *Broker) Subscribe(topic string, conn net.Conn) {
	b.Subscribers[topic] = append(b.Subscribers[topic], conn)
	fmt.Printf("New subscriber added to topic: %s\n", topic)
}
