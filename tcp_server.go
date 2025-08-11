package main

import (
	"log"
	"net"
)

func main() {
	const (
		HOST = "localhost"
		PORT = "8080"
		TYPE = "tcp"
	)

	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Awaiting TCP connection on port " + PORT)

	conn, err := listener.Accept()
	if err != nil {
		log.Println("Listener error occurred.")
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("Connection established!")
	log.Println("Remote endpoint address: " + conn.RemoteAddr().String())

	buffer := make([]byte, 1024)
	
	for {
		log.Println("Awaiting incoming message (1024b buffer)..")
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("TCP Conn error: " + err.Error())
			break
		}
		log.Println("Message data: " + string(buffer[:n]))
	}
}
