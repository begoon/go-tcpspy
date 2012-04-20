// TCP/IP debugger/proxy
// Copyright (C) 2012 by Alexander Demin

package main

import (
  "encoding/hex"
  "flag"
  "fmt"
  "net"
  "os"
  "runtime"
  "strings"
  "time"
)

var (
  host        *string = flag.String("host", "", "target host or address")
  port        *string = flag.String("port", "0", "target port")
  listen_port *string = flag.String("listen_port", "0", "listen port")
)

func die(format string, v ...interface{}) {
  os.Stderr.WriteString(fmt.Sprintf(format+"\n", v...))
  os.Exit(1)
}

func connection_logger(data chan string, conn_n int, local_info, remote_info string) {
  log_name := fmt.Sprintf("log-%s-%04d-%s-%s.log", format_time(time.Now()),
    conn_n, local_info, remote_info)
  f, err := os.Create(log_name)
  if err != nil {
    die("Unable to create file %s, %v\n", log_name, err)
  }
  defer f.Close()
  for {
    s := <-data
    if s == "" {
      break
    }
    f.WriteString(s)
    f.Sync()
  }
}

func binary_logger(data chan []byte, conn_n int, peer string) {
  log_name := fmt.Sprintf("log-binary-%s-%04d-%s.log", format_time(time.Now()),
    conn_n, peer)
  f, err := os.Create(log_name)
  if err != nil {
    die("Unable to create file %s, %v\n", log_name, err)
  }
  defer f.Close()
  for {
    b := <-data
    if len(b) == 0 {
      break
    }
    f.Write(b)
    f.Sync()
  }
}

func format_time(t time.Time) string {
  return t.Format("2006.01.02-15.04.05")
}

func printable_addr(a net.Addr) string {
  return strings.Replace(a.String(), ":", "-", -1)
}

type Channel struct {
  from, to      net.Conn
  logger        chan string
  binary_logger chan []byte
  ack           chan bool
}

func pass_through(c *Channel) {
  from_peer := printable_addr(c.from.LocalAddr())
  to_peer := printable_addr(c.to.LocalAddr())

  b := make([]byte, 10240)
  offset := 0
  packet_n := 0
  for {
    n, err := c.from.Read(b)
    if err != nil {
      c.logger <- fmt.Sprintf("Disconnected from %s\n", from_peer)
      break
    }
    if n > 0 {
      c.logger <- fmt.Sprintf("Received (#%d, %08X) %d bytes from %s\n",
        packet_n, offset, n, from_peer)
      c.logger <- hex.Dump(b[:n])
      c.binary_logger <- b[:n]
      c.to.Write(b[:n])
      c.logger <- fmt.Sprintf("Sent (#%d) to %s\n", packet_n, to_peer)
      offset += n
      packet_n += 1
    }
  }
  c.from.Close()
  c.to.Close()
  c.ack <- true
}

func process_connection(local net.Conn, conn_n int, target string) {
  remote, err := net.Dial("tcp", target)
  if err != nil {
    fmt.Printf("Unable to connect to %s, %v\n", target, err)
  }

  local_info := printable_addr(remote.LocalAddr())
  remote_info := printable_addr(remote.RemoteAddr())

  started := time.Now()

  logger := make(chan string)
  from_logger := make(chan []byte)
  to_logger := make(chan []byte)
  ack := make(chan bool)

  go connection_logger(logger, conn_n, local_info, remote_info)
  go binary_logger(from_logger, conn_n, local_info)
  go binary_logger(to_logger, conn_n, remote_info)

  logger <- fmt.Sprintf("Connected to %s at %s\n", target, format_time(started))

  go pass_through(&Channel{remote, local, logger, to_logger, ack})
  go pass_through(&Channel{local, remote, logger, from_logger, ack})
  <-ack // Make sure that the both copiers gracefully finish.
  <-ack // 

  finished := time.Now()
  duration := finished.Sub(started)
  logger <- fmt.Sprintf("Finished at %s, duration %s\n", format_time(started), duration.String())

  logger <- ""            // Stop logger
  from_logger <- []byte{} // Stop "from" binary logger
  to_logger <- []byte{}   // Stop "to" binary logger
}

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())
  flag.Parse()
  if flag.NFlag() != 3 {
    fmt.Printf("usage: gotcpspy -host target_host -port target_port -listen_post=local_port\n")
    flag.PrintDefaults()
    os.Exit(1)
  }
  target := net.JoinHostPort(*host, *port)
  fmt.Printf("Start listening on port %s and forwarding data to %s\n", *listen_port, target)
  ln, err := net.Listen("tcp", ":"+*listen_port)
  if err != nil {
    fmt.Printf("Unable to start listener, %v\n", err)
    os.Exit(1)
  }
  conn_n := 1
  for {
    if conn, err := ln.Accept(); err == nil {
      go process_connection(conn, conn_n, target)
      conn_n += 1
    } else {
      fmt.Printf("Accept failed, %v\n", err)
    }
  }
}
