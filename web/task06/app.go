package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn){
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan(){
		log.Println(scanner.Text())
		if scanner.Text() == ""{
			log.Println("End of the HTTP header")
			break
		}
	}

	fmt.Println("Code got here")
	io.WriteString(conn, "I see you connected.")
}
