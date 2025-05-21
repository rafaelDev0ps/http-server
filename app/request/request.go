package request

import (
	"strings"
)

type Request struct {
	Method         string
	Path           string
	Version        string
	RequestHeaders RequestHeaders
	RequestBody    []byte
}

type RequestHeaders []string

// TODO: create a request parser to parse the request

func (req *Request) GetMethod(reqData []string) string {
	return strings.Split(reqData[0], " ")[0]
}

func (req *Request) GetPath(reqData []string) string {
	return strings.Split(reqData[0], " ")[1]
}

func (req *Request) GetProtocolVersion(reqData []string) string {
	return strings.Split(reqData[0], " ")[2]
}

func (req *Request) GetHeaders(reqData []string) RequestHeaders {
	return reqData[1 : len(reqData)-2]
}

func (req *Request) GetHeaderValue(headers []string, fieldName string) string {
	var acceptEncoding string
	for i := range len(headers) {
		if strings.Contains(headers[i], fieldName) {
			acceptEncoding = headers[i]
			return strings.Split(acceptEncoding, ": ")[1]
		}
	}

	return ""
}

func (req *Request) GetBody(reqData []string) string {
	return reqData[len(reqData)-1]
}
