package main

import (
  "flag"
  "fmt"
  "net"
  "os"
)

var (
  host        *string = flag.String("host", "", "target host or address")
  port        *string = flag.String("port", "0", "target port")
  listen_port *string = flag.String("listen_port", "0", "listen port")

  target string
  conn_n int
)

func process_connection(local net.Conn) {
  target, err := net.Dial("tcp", *host+":"+*port)
  if err != nil {
    fmt.Printf("Unable to connect to %s, %v\n", target, err)
  }
  target.Close()
  local.Close()
}

func main() {
  flag.Parse()
  if flag.NFlag() != 3 {
    flag.PrintDefaults()
    os.Exit(1)
  }
  target = net.JoinHostPort(*host, *port)
  fmt.Printf("Start listening on port %s and forwarding data to %s\n", *listen_port, target)
  ln, err := net.Listen("tcp", ":"+*listen_port)
  if err != nil {
    fmt.Printf("Unable to start listener, %v\n", err)
    os.Exit(1)
  }
  for {
    if conn, err := ln.Accept(); err == nil {
      fmt.Printf("%v\n", conn)
      go process_connection(conn)
    } else {
      fmt.Printf("Accept failed, %v\n", err)
    }
  }
}
