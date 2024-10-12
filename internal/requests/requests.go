package requests

import (
	"errors"
	"strconv"
	"strings"
)

type RMethod string

const (
	GET     RMethod = "GET"
	HEAD    RMethod = "HEAD"
	POST    RMethod = "POST"
	PUT     RMethod = "PUT"
	DELETE  RMethod = "DELETE"
	TRACE   RMethod = "TRACE"
	CONNECT RMethod = "CONNECT"
	INVALID RMethod = "INVALID"
)

type RHeader string

type Request struct {
	Method       RMethod
	ResourcePath string
	HttpVersion  float32

	Headers []RHeader
}

func New(requestLines []string) (*Request, error) {

	rOverview := strings.Fields(requestLines[0])

	if len(rOverview) < 3 || len(rOverview) > 3 {

		return nil, errors.New("Expected only 3 elements on request overview")
	}

	method := getMethodBasedOnText(rOverview[0])
	path := getPathBaseOnText(rOverview[1])
	httpVersion, errVersion := getHttpVersionBasedOnText(rOverview[2])

	if method == INVALID ||
		errVersion != nil {

		return nil, errors.New("Error while trying to create request object")
	}

	return &Request{
		Method:       method,
		ResourcePath: path,
		HttpVersion:  httpVersion,
	}, nil
}

func getMethodBasedOnText(rMethod string) RMethod {

	methods := map[string]RMethod{
		"GET":     GET,
		"HEAD":    HEAD,
		"POST":    POST,
		"PUT":     PUT,
		"DELETE":  DELETE,
		"TRACE":   TRACE,
		"CONNECT": CONNECT,
    }
    
    value, ok := methods[rMethod]
    if ok {
        return value
    } 
    return INVALID
}

func getPathBaseOnText(rPath string) string {

	const ROOT = "./www/"

	if rPath == "" {
		rPath = "/"
	}

	rPath = ROOT + rPath[1:]

	if rPath == ROOT {
		return rPath + "index.html"
	}

	return rPath
}

func getHttpVersionBasedOnText(httpVersion string) (float32, error) {

	versionInText := strings.Replace(httpVersion, "HTTP/", "", -1)
	version, err := strconv.ParseFloat(versionInText, 32)

	if err != nil {
		return 0.0, errors.New("Cannot get appropiate version of the request")
	}

	return float32(version), nil
}
