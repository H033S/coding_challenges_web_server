package core

import (
	"bytes"
	"errors"
	"net"
	"strings"

	"github.com/H033S/web_server/internal/requests"
)

const CHUNK_SIZE = 4096

func HandleRequest(conn *net.Conn) (*requests.Request, error) {

	received, buffer, err := getRawRequestInformationInBytes(conn)

	if err != nil {
		return nil, errors.New("Error detected while trying to retrieve information from request\n")
	}
	if received <= 0 {
		return nil, errors.New("Request cannot be blank. Erroing out!!!\n")
	}

	requestInPlainText := string(buffer)
	requestSplittedPerLine := strings.Split(requestInPlainText, "\r\n")

	return requests.New(requestSplittedPerLine)
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
