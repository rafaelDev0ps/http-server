package controller

import (
	"fmt"
	"net"

	"http-server/app/request"
	"http-server/app/response"
)

func DefaultController(conn net.Conn, content []string) (int, error) {
	var res response.Response
	var req request.Request
	res.ResponseHeader = make(response.ResponseHeader)

	headers := req.GetHeaders(content)

	if req.GetHeaderValue(headers, "Connection") == "close" {
		res.AddHeader("Connection", "close")
	}
	return fmt.Fprintf(conn, "%s%s\r\n", res.Status200(), res.FormatHeaders())
}
