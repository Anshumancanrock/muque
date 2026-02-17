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

type Subscriber struct {
	Conn net.Conn
}

type Broker struct {
	// Topic -> Array of Subscribers (what I call the hash table of pub/sub)
	Subscribers map[string][]*Subscriber
}

func (b *Broker) Subscribe(topic string, sub *Subscriber) {
	b.Subscribers[topic] = append(b.Subscribers[topic], sub)
	fmt.Printf("New subscriber added to topic: %s\n", topic)
}

func (b *Broker) Publish(topic string, payload string) {
	subscribers := b.Subscribers[topic]
	outgoingMsg := []byte(fmt.Sprintf("BROKER BROADCAST [%s]: %s\n", topic, payload))

	for _, sub := range subscribers {
		sub.Conn.Write(outgoingMsg)
	}

	fmt.Printf("Broadcasted to %d subscribers on topic: %s\n", len(subscribers), topic)
}
