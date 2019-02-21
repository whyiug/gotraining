package main

import (
	"fmt"
	"net"
	"os"
)

func sender(conn net.Conn) int {
	var input string
	username := conn.LocalAddr().String()
	lens, err := conn.Write([]byte("forrest\r\n"))
	for {
		fmt.Scanln(&input)
		// fmt.Println([]byte input)

		lens, err := conn.Write([]byte(username + " Say :::" + input))
		fmt.Println(lens)
		defer conn.Close()
		if err != nil {
			fmt.Println(err.Error())
			break
		}

	}

}

func main() {
	server := "192.168.0.141:12345"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	for true {
		sender(conn)
	}

}
