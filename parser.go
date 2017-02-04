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

	GET_USAGE = "Invalid use of get:\nUsage: get [key]"
	ADD_USAGE = "Invalid usage of add:\nUsage: add [key] [value]"
)

func ParseCommand(s string) string {
	if len(s) == 0 {
		return ""
	}

	words := strings.Split(s, " ")
	cmd := words[0]
	datastore := GetStore()
	switch cmd {
	case GET:
		if len(words) != 2 {
			return GET_USAGE
		}
		key := words[1]
		result := datastore.Get(key) + "\n"
		return result
	case ADD:
		if len(words) < 3 {
			return ADD_USAGE
		}
		key := words[1]
		value := words[2]
		datastore.Add(key, value)
		return "Added!\n"
	case DELETE:
		datastore.Delete(words[1])
		return "Deleted!\n"
	case DUMP:
		fmt.Println("Sending dump")
		return datastore.Dump() + "\n"
	case STATUS:
		return "OK\n"
	default:
		return "Unrecognized command\n"
	}
}
