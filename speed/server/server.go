package main

import (
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":49152")
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