package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var peers *list.List

func queryPeer(addr string) {
	// TODO: send a message to peer to
	// return all k-v data serialized
	ln, err := net.Listen("tcp", "5000")
	defer ln.Close()

	if err != nil {
		fmt.Printf("Couldn't connect to peer %s\n", addr)
	}

	conn, err := ln.Accept()

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(line)
	conn.Close()
}

func collectPeers() {
	peers = list.New()
	// parse arguments
	if len(os.Args) > 1 {
		fmt.Println("There are peers")
	}

	for i := 0; i < len(os.Args)-1; i += 1 {
		peerAddr := os.Args[i]
		queryPeer(peerAddr)
		peers.PushBack(peerAddr)
	}
}

func socketHandler(c net.Conn) {
	for {
		reader := bufio.NewReader(c)
		writer := bufio.NewWriter(c)

		line, _ := reader.ReadString('\n')
		result := ParseCommand(strings.TrimSpace(line))

		writer.WriteString(result)
		writer.Flush()
	}
	// Shut down the connection
	c.Close()
}

func main() {
	collectPeers()
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go socketHandler(conn)
	}
}
