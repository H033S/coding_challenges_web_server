package responses

import "github.com/H033S/web_server/internal/requests"

func CreateNotFoundResponse(rq *requests.Request) *Response {

    return &Response{
        statusCode: 404,
        statusDescription: "Not Found",
        content: make([]byte, 0),
    }
}
