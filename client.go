package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	if response != "OK\n" {
		fmt.Print("Bad response", response)
	} else {
		fmt.Print("OK")
	}

	conn.Close()
}
