package main

import (
	"bufio"
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var peers *list.List

func queryPeer(addr string) {

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("Couldn't connect to peer %s\n", addr)
	}
	defer conn.Close()
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	writer.WriteString(DUMP + "\n")
	writer.Flush()

	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(line)

	var m map[string]string
	err = json.Unmarshal([]byte(line), &m)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range m {
		GetStore().Add(k, v)
	}
}

func collectPeers() {
	peers = list.New()

	for i := 2; i < len(os.Args); i += 1 {
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

	ln, err := net.Listen("tcp", os.Args[1])
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
