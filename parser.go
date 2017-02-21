package main

import (
	"fmt"
	"strings"
)

const (
	GET    = "get"
	ADD    = "add"
	DELETE = "delete"
	DUMP   = "dump"
	STATUS = "status"
	PEER   = "peer"
)

func ParsePeerCommand(s []string) string {
	switch s[1] {
	case ADD:
		peerIP := s[2]
		queryPeer(peerIP)
		peers.PushBack(peerIP)
		return "OK\n"
	default:
		return "Unrecognized command\n"
	}
}

func ParseCommand(s string) string {
	if len(s) == 0 {
		return ""
	}

	words := strings.Split(s, " ")
	cmd := words[0]
	datastore := GetStore()
	switch cmd {
	case GET:
		return GetCmd(words)
	case ADD:
		return AddCmd(words)
	case DELETE:
		return DeleteCmd(words)
	case DUMP:
		return datastore.Dump() + "\n"
	case STATUS:
		return "OK\n"
	case PEER:
		return ParsePeerCommand(words)
	default:
		return "Unrecognized command\n"
	}
}
