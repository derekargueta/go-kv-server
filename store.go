package main

import (
	"encoding/json"
	"fmt"
)

type store struct {
	table map[string]string
}

var instance *store

func GetStore() *store {
	if instance == nil {
		fmt.Println("Creating a new datastore")
		tmpMap := make(map[string]string)
		instance = &store{tmpMap}
	}

	return instance
}

func (s *store) Add(k, v string) {
	s.table[k] = v
	fmt.Printf("Added %s: %s\n", k, v)
}

func (s *store) Get(k string) string {
	// todo: error handling (key error)
	fmt.Printf("Fetching %s\n", k)
	return s.table[k]
}

func (s *store) Delete(k string) {
	delete(s.table, k)
}

func (s *store) Dump() string {
	jsonBytes, _ := json.Marshal(s.table)
	return string(jsonBytes)
}
