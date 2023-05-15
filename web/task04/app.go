package main

import (
	"fmt"
	"net"
	"io"
)
func main() {

	 listener, error := net.Listen("tcp",":8080")
	 if error != nil {
		fmt.Println(error)
	 }
	 defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	io.WriteString(conn, "I see you connected")

}