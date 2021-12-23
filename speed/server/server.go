package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var (
	port = flag.String("port", "49152", "listen port")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf(":%s", *port)
	log.Printf("listen on %s\n", addr)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(c)
	}
}

func handle(c net.Conn) {
	defer c.Close()

	buf := make([]byte, 1024)
	res := make([]byte, 1024)

	n, err := c.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}

	copy(res, buf[:n])

	n, err = c.Write(res[:n])
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("hold")
	time.Sleep(5*time.Second)
	log.Println("stop")
}