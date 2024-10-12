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

		go internal.HandleRequest(&conn)
	}
}
