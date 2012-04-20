package main

import (
  "flag"
  "fmt"
  "net"
  "os"
  "strings"
  "time"
)

var (
  host        *string = flag.String("host", "", "target host or address")
  port        *string = flag.String("port", "0", "target port")
  listen_port *string = flag.String("listen_port", "0", "listen port")

  target string
)

func die(format string, v ...interface{}) {
  os.Stderr.WriteString(fmt.Sprintf(format+"\n", v...))
  os.Exit(1)
}

func connection_logger(log_name string, data chan string) {
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
    println(s)
    f.WriteString(s)
    f.WriteString("\n")
    f.Sync()
  }
}

func format_time(t time.Time) string {
  return t.Format("2006.01.02-15.04.05")
}

func transfer(from, to net.Conn, logger chan string, ack chan bool, info string) {
  b := make([]byte, 10240)
  for {
    n, err := from.Read(b)
    if err != nil {
      logger <- fmt.Sprintf("Disconnected from %s", info)
      break
    }
    if n > 0 {
      logger <- fmt.Sprintf("[%s]", string(b[:n]))
      to.Write(b[:n])
    }
  }
  from.Close()
  to.Close()
  ack <- true
}

func process_connection(local net.Conn, conn_n int) {
  remote, err := net.Dial("tcp", target)
  if err != nil {
    fmt.Printf("Unable to connect to %s, %v\n", target, err)
  }
  
  local_info := strings.Replace(remote.LocalAddr().String(), ":", "-", -1)
  remote_info := strings.Replace(remote.RemoteAddr().String(), ":", "-", -1)
  
  started := time.Now()
  log_name := fmt.Sprintf("log-%s-%04d-%s-%s.log", format_time(started), conn_n, local_info, remote_info)
  
  logger := make(chan string)
  ack := make(chan bool)
  
  go connection_logger(log_name, logger)

  logger <- fmt.Sprintf("Connected to %s at %s", target, format_time(started))

  go transfer(remote, local, logger, ack, remote_info)
  go transfer(local, remote, logger, ack, local_info)
  <-ack
  <-ack

  finished := time.Now()
  duration := finished.Sub(started)
  logger <- fmt.Sprintf("Finished at %s, duration %s", format_time(started), duration.String())

  logger <- ""
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
  conn_n := 1
  for {
    if conn, err := ln.Accept(); err == nil {
      go process_connection(conn, conn_n)
    } else {
      fmt.Printf("Accept failed, %v\n", err)
    }
  }
}
