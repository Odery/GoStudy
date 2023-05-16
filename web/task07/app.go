package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	var request []string

	for scanner.Scan() {
		v := scanner.Text()
		log.Println(v)
		request = append(request, v)
		if v == "" {
			log.Println("End of the HTTP header")
			break
		}
	}

	//We dont know what piece of application will acccess this TCP 8080 port
	//It could have no HTTP headers and so, without this check, our program will crash
	if len(request) != 0 {
		body := muxHandler(request)
		io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")

		io.WriteString(conn, "\r\n")
		io.WriteString(conn, body)
	}
}

type request struct {
	method string
	path string
}


//This is a mux server for a current application
//I know this can be done better, but for simplicity
//this solution is better.
//This method only returns different html, depending on
//method and path requested, again, better for simplicity
func muxHandler(r []string) string{
	req := request{method: strings.Split(r[0], " ")[0],
				   path: strings.Split(r[0], " ")[1]}
	var msg string

	if req.method == "GET" && req.path == "/"{
		msg = "You have requested main page via GET request"
	}else if req.method == "GET" && req.path == "/apply"{
		msg = "You have requested apply page via GET request"
	}else if req.method =="POST" && req.path == "/apply"{
		msg = "You have POSTed some data to the /apply page!"
	}else{
		msg = "404 - Resource not found"
	}

	body := fmt.Sprintf(`<html>
								<head><title>Main page</title>
								</head>
							<body>
								<h1>Hello there traveller</h1>
								<h2>%v</h2>
								<ul>
									<li>Request method: %v</li>
									<li>Request URI: %v</li>
								</ul>
							</body>
							</html>`, msg, req.method, req.path)

	return body
}
