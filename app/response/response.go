package response

import (
	"fmt"
	"net"
)

type Response struct {
	StatusCode      string
	ProtocolVersion string
	ResponseHeader  ResponseHeader
}

type ResponseHeader map[string]string

func (res *Response) AddHeader(key string, value string) {
	res.ResponseHeader[key] = value
}

func (res *Response) FormatHeaders() string {
	var headersString string
	for key, header := range res.ResponseHeader {
		headersString += key + ": " + header + "\r\n"
	}
	return headersString
}

func (res *Response) GetResponseLine() string {
	return res.ProtocolVersion + " " + res.StatusCode + "\r\n"
}

func (res *Response) Status200() string {
	r := Response{
		StatusCode:      "200 OK",
		ProtocolVersion: "HTTP/1.1",
	}
	return r.ProtocolVersion + " " + r.StatusCode + "\r\n"
}

func (res *Response) Status500(conn net.Conn) (int, error) {
	resp := &Response{
		StatusCode:      "500 Internal Server Error",
		ProtocolVersion: "HTTP/1.1",
	}
	return fmt.Fprintf(conn, "%s\r\n", resp.GetResponseLine())
}

func (res *Response) Status404(conn net.Conn) (int, error) {
	resp := &Response{
		StatusCode:      "404 Not Found",
		ProtocolVersion: "HTTP/1.1",
	}
	return fmt.Fprintf(conn, "%s\r\n", resp.GetResponseLine())
}
