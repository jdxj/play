package main

import (
	"flag"
	"log"
	"net"
	"time"
)

var (
	addr = flag.String("addr", "", "remote addr")
)

func main() {
	flag.Parse()
	if *addr == "" {
		log.Fatalln("addr empty")
	}

	buf := make([]byte, 1024)

	t1 := time.Now()

	c, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	t2 := time.Now()

	_, err = c.Write([]byte("hello world"))
	if err != nil {
		log.Fatalln(err)
	}

	t3 := time.Now()

	n, err := c.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	t4 := time.Now()

	log.Println(buf[:n])
	log.Printf("dial: %d\n", t2.Sub(t1).Milliseconds())
	log.Printf("write: %d\n", t3.Sub(t2).Milliseconds())
	log.Printf("read: %d\n", t4.Sub(t3).Milliseconds())
	log.Printf("write+read: %d\n", t4.Sub(t2).Milliseconds())
	log.Printf("all: %d\n", t4.Sub(t1).Milliseconds())
}
