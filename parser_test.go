package main

import (
	"testing"
)

func TestParsingAdd(t *testing.T) {
	result := ParseCommand("add hi hello")
	if result != "Added!\n" {
		t.Fail()
	}

	result = ParseCommand("add hi")
	if result != ADD_USAGE {
		t.Fail()
	}

	result = ParseCommand("add")
	if result != ADD_USAGE {
		t.Fail()
	}
}

func TestParsingGet(t *testing.T) {
	// stick something in the store to retrieve as a test
	GetStore().Add("hi", "hello")
	result := ParseCommand("get hi")
	if result != "hello\n" {
		t.Fail()
	}

	result = ParseCommand("get")
	if result != GET_USAGE {
		t.Fail()
	}

	result = ParseCommand("get hi hello hey")
	if result != GET_USAGE {
		t.Fail()
	}
}
