package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	//Infinete loop for wating clients
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("Connection Established")

		// write to connection
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")
		// run command in your terminal : telnet localhost 8080

		conn.Close()
	}
}
