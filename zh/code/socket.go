package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept err:", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	bytesRead, error := conn.Read(buffer)
	if error != nil {
		log.Println("Client connection error: ", error)
	}
	log.Println("get content from client:", string(buffer[0:bytesRead]))
}
