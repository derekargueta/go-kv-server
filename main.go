
package main

import (
  "net"
  "bufio"
  "log"
)

func main() {
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
    go func(c net.Conn) {
      reader := bufio.NewReader(c)
      writer := bufio.NewWriter(c)

      line, _ := reader.ReadString('\n')
      result := ParseCommand(line)

      writer.WriteString(result)
      writer.Flush()
      // Shut down the connection
      c.Close()
    }(conn)
  }
}
