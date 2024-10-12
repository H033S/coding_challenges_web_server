package main

import (
	"fmt"
	"net"
	"os"

	internal "github.com/H033S/web_server/internal"
)

func main() {

	ln, err := net.Listen("tcp", ":80")
	if err != nil {

		fmt.Println("An error has ocurred")
		os.Exit(0)
	}

	for {

		conn, err := ln.Accept()
		if err != nil {

			fmt.Println("An error has ocurred")
			os.Exit(0)
		}

		rq, err := internal.HandleRequest(&conn)

		if err != nil {
			fmt.Println("Something went wrong while creating request")
			os.Exit(0)
		}

		content, err := os.ReadFile((*rq).ResourcePath)
		response := fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\n%s\r\n", content)
		responseInBytes := []byte(response)

        fmt.Println(content)

		conn.Write(responseInBytes)
		conn.Close()
	}
}
