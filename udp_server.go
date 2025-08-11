package main

import (
	"log"
	"net"
	"time"
	"fmt"
	"os"
)

func main() {
	localAddr := os.Args[1]
	udpServer, err := net.ListenPacket("udp", localAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()
	
	log.Println("Now listening for packets on port 8080..")

	for {
		buf := make([]byte, 1024)
		n, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			log.Println("ReadFrom error: " + err.Error())
			break
		}
		log.Println("Received: " + string(buf[:n]) + " from " + addr.String())
		go response(udpServer, addr, buf[:n])
	}
}

func response(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("time received: %v. Your message: %v!", time, string(buf))

	udpServer.WriteTo([]byte(responseStr), addr)
}
