package main

import (
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
