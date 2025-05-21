package request

import (
	"strings"
)

type Request struct {
	Method          string
	Path            string
	ProtocolVersion string
	RequestHeaders  RequestHeaders
	RequestBody     []byte
}

type RequestHeaders map[string]string

func (req *Request) ParseRequest(reqData []string) Request {
	var request Request
	request.Method = strings.Split(reqData[0], " ")[0]
	request.Path = strings.Split(reqData[0], " ")[1]
	request.ProtocolVersion = strings.Split(reqData[0], " ")[2]

	request.RequestHeaders = make(RequestHeaders)

	rawHeaders := reqData[1:]
	if request.Method != "GET" {
		rawHeaders = reqData[1 : len(reqData)-2]
	}

	if len(rawHeaders) > 0 {
		for _, headerLine := range rawHeaders {
			keyValue := strings.Split(headerLine, ": ")
			if len(keyValue) > 1 {
				request.RequestHeaders[keyValue[0]] = keyValue[1]
			}
		}
	}

	if request.Method != "GET" {
		bodyString := reqData[len(reqData)-1] //get last line
		if bodyString != "" {
			request.RequestBody = []byte(bodyString)
		}
	}

	return request
}
