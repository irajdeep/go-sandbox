package main

import (
	"io"
	"log"
	"net"
	"os"
)

const (
	sockAddr = "/tmp/echo.sock"
)

func echoServer(c net.Conn) {
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
	io.Copy(c, c)
	c.Close()
}

func main() {
	if err := os.RemoveAll(sockAddr); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", sockAddr)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error: ", err)
		}
		go echoServer(conn)

	}

}
