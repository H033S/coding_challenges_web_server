package core

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/H033S/web_server/internal/requests"
	"github.com/H033S/web_server/internal/responses"
)

const CHUNK_SIZE = 4096
func HandleRequest(conn *net.Conn){

	received, buffer, err := getRawRequestInformationInBytes(conn)

	if err != nil {
        
		fmt.Println("Error detected while trying to retrieve information from request")
        os.Exit(0)
	}
	if received <= 0 {

		fmt.Println("Request cannot be blank. Erroing out!!!")
        os.Exit(0)
	}

	requestInPlainText := string(buffer)
	requestSplittedPerLine := strings.Split(requestInPlainText, "\r\n")

    rq, err := requests.New(requestSplittedPerLine)
    if err != nil {
        
        fmt.Println("Bad Request")
        os.Exit(0)
    }

    rp, err := responses.New(rq)
    if err != nil {

        fmt.Println("Bad Response")
        os.Exit(0)
    }
    time.Sleep(10 * time.Second)
    fmt.Println("Writing to console")
    rp.Write(conn)
    (*conn).Close()
}

func getRawRequestInformationInBytes(conn *net.Conn) (int, []byte, error) {

	var received int
	buffer := bytes.NewBuffer(nil)

	for {

		chunk := make([]byte, CHUNK_SIZE)
		read, err := (*conn).Read(chunk)
		if err != nil {
			return received, buffer.Bytes(), err
		}

		received += read
		buffer.Write(chunk[:read])

		if read == 0 || read < CHUNK_SIZE {
			break
		}
	}

	return received, buffer.Bytes(), nil
}
