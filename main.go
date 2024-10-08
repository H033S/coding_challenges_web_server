package main

import (
	"fmt"
	"net"
	"strings"
)

func handleRequest(conn net.Conn) {

	var buffer = make([]byte, 1024)
	cp, err := conn.Read(buffer)
	if err != nil {

		fmt.Print("An error has ocurred")
	}

	if cp > 0 {

        fmt.Print("Parsing path in request")
		var path string = retrievePath(buffer)

        fmt.Print("Generating response")
        response := fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\nRequested Path: %v \r\n", path)

        fmt.Print("Retrieving response back to server")
        conn.Write([]byte(response))
	}
}

func retrievePath(request []byte) string {
    
    sRequest := string(request)
    result := strings.Split(sRequest, "\r\n")
    if len(result) > 0 {

        requestLine := strings.Fields(result[0])
        return requestLine[1]
    }

    return ""
}

func startListeningFromConnections() {

	ln, err := net.Listen("tcp", ":80")
	if err != nil {

		fmt.Print("An error has ocurred")
	}

	for {

		conn, err := ln.Accept()
		if err != nil {

			fmt.Print("An error has ocurred")
		}

        fmt.Print("Connection accepted")
		handleRequest(conn)
        conn.Close()
	}
}

func main() {

	startListeningFromConnections()
}
