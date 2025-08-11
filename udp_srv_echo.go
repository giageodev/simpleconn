package main

import (
	"log"
	"net"
	"strconv"
	"os"
)

func main() {
	localAddr := os.Args[1]
	udpServer, err := net.ListenPacket("udp", localAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()
	
	log.Println("Now listening for packets on " + localAddr)

	for i:=0; i<1000000; i++ {
		buf := make([]byte, 1024)
		n, _, err := udpServer.ReadFrom(buf)
		if err != nil {
			log.Println("ReadFrom error: " + err.Error())
			os.Exit(1)
		}

		num, err := strconv.Atoi(string(buf[:n]))
		if err != nil {
			log.Fatal(err)
		}

		if num != i {
			log.Printf("Loop counter does not match packet counter! @ i=%v", i)
			log.Printf("Packet content was: %v", buf[:n])
			os.Exit(1)
		}
	}
	
	log.Println("Received 1 million UDP packets, uncorrupted and in order!")
}
