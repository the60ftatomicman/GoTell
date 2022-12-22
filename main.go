package main

import (
	"example/gotell/src/session"
	"fmt"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "9002"
	TYPE = "tcp"
)

// stty -icanon && nc localhost 9002
func main() {
	log.Println("Connect via stty -icanon && nc " + HOST + " " + PORT)
	listner, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listner.Close()

	for {
		c, err := listner.Accept()
		log.Println("looks like we have a connection!")
		if err != nil {
			fmt.Println(err)
			return
		}
		s := session.Session{}
		s.Initialize(&c)
		//Now begin handling
		go s.Handle() //handleConnection(c, p, s, m)
	}
}
