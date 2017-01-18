
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
  if result != AddUsage() {
    t.Fail()
  }

  result = ParseCommand("add")
  if result != AddUsage() {
    t.Fail()
  }
}
