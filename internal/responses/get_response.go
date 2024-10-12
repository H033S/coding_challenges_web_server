package responses

import (
	"fmt"
	"os"

	"github.com/H033S/web_server/internal/requests"
)

func CreateGetResponse(rq *requests.Request) *Response {

    fmt.Println("Creating request")
    content, err := os.ReadFile(rq.ResourcePath)
    if err != nil {
        return CreateNotFoundResponse(rq)
    }
    
    return &Response{

        statusCode: 200,
        statusDescription: "OK",
        content: content,
    }
}


func validPath (path string) bool {

    _, err := os.Stat(path)
    if err != nil {

        return false
    }

    return true
}
