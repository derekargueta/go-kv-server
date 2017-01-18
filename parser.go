package main

import (
  "strings"
)

const (
  GET = "get"
  ADD = "add"
  DELETE = "delete"
)


func AddUsage() string {
  return "Invalid usage of add:\nUsage: add [key] [value]"
}


func ParseCommand(s string) string {
  words := strings.Split(s, " ")
  cmd := words[0]
  datastore := GetStore()
  switch cmd {
  case GET:
    return datastore.Get(words[1]) + "\n"
  case ADD:
    if len(words) < 3 {
      return AddUsage()
    }
    key := words[1]
    value := words[2]
    datastore.Add(key, value)
    return "Added!\n"
  case DELETE:
    datastore.Delete(words[1])
    return "Deleted!\n"
  default:
    return "Unrecognized command\n"
  }
}
