package main

import (
  "strings"
  "fmt"
)

const (
  GET = "get"
  ADD = "add"
  DELETE = "delete"
)


func ParseCommand(s string) string {
  words := strings.Split(s, " ")
  cmd := words[0]
  datastore := GetStore()
  switch cmd {
  case GET:
    fmt.Println("Handling get with key: " + words[1])
    return datastore.Get(words[1]) + "\n"
  case ADD:
    key := words[1]
    value := words[2]
    fmt.Println("Handling add with key " + key + " and value " + value)
    datastore.Add(key, value)
    return "Added!\n"
  case DELETE:
    fmt.Println("Handling delete")
    datastore.Delete(words[1])
    return "Deleted!\n"
  default:
    return "Unrecognized command\n"
  }
}
