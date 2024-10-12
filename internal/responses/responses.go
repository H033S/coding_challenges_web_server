package responses

import (
	"errors"
	"net"
	"strconv"

	"github.com/H033S/web_server/internal/requests"
)

type Response struct {
	statusCode        int
	statusDescription string
	content           []byte
}

func New(rq *requests.Request) (*Response, error) {

	methods := map[requests.RMethod]func(*requests.Request) *Response{
		requests.GET: CreateGetResponse,
	}

	method, ok := methods[rq.Method]
	if ok {

		return method(rq), nil
	}
	return nil, errors.New("Invalid or not Implemented Method")
}

func (rp *Response) Write(conn *net.Conn) {

	message := rp.GetMessageInBytes()
	(*conn).Write(message)
}

func (rp *Response) GetMessageInBytes() []byte {

	message := "HTTP/1.1 " + strconv.Itoa(rp.statusCode) + " " + rp.statusDescription + "\r\n"
	message += "\r\n"
	message += string(rp.content)
	message += "\r\n"

	return []byte(message)
}
