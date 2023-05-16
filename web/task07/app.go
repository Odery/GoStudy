package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"fmt"
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

	body := "<html><head><title>Main page</title></head><body><p>Hello there traveller</p></body></html>"

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")

	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
