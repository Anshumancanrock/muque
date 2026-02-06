package main

type Message struct {
	Command string `json:"command"`
	Topic   string `json:"topic"`
	Payload string `json:"payload"`
}
