package main

import (
	"io"
	"log"
	"net"
)

var listenAddr = ":8080"

var forwardAddr ="example.com:80"

func handleConnection(client net.Conn){
	defer client.Close()

	server, err := net.Dial("tcp", forwardAddr)
	if err != nil{
		log.Println("error server remote", err)
		return
	}

	defer server.Close()
	log.Printf("🔁 Forwarding: %s <--> %s\n", client.RemoteAddr(), forwardAddr)

	// Copiază datele bidirecțional (client <-> server)
	go io.Copy(server, client)
	io.Copy(client, server)
}

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalln("❌ Eroare la ascultare:", err)
	}
	log.Println("🚪 Ascult pe", listenAddr, "->", forwardAddr)

	for {
		client, err := listener.Accept()
		if err != nil {
			log.Println("❌ Eroare acceptare conexiune:", err)
			continue
		}
		go handleConnection(client)
	}
}