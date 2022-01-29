package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error:%v\n", err)

	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error:%v\n", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			log.Printf("read error:%v\n", err)
			return
		}
		wr.WriteString("hello ")
		wr.Write(line)
		wr.Flush()
	}
}
