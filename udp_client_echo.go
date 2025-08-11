package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	conn, err := net.Dial("udp", os.Args[1])
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	for i:=0; i < 1000000; i++ {
		_, err := conn.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			log.Printf("Error sending packet #%v: %v\n", i, err.Error())
			os.Exit(1)
		} 
		time.Sleep(time.Millisecond * 100)
		 
		if i % 100 == 0 {
			log.Printf("Currently at packet %v\n", i)
		}
	}

	log.Printf("Sent 1 million UDP packets to %v", os.Args[1])
}
