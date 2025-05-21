package controller

import (
	"fmt"
	"net"

	"http-server/app/request"
	"http-server/app/response"
)

// Return response object instead (int, error)
func DefaultController(conn net.Conn, request request.Request) (int, error) {
	var res response.Response
	res.ResponseHeader = make(response.ResponseHeader)

	if request.RequestHeaders["Connection"] == "close" {
		res.AddHeader("Connection", "close")
	}

	return fmt.Fprintf(conn, "%s%s\r\n", res.Status200(), res.FormatHeaders())
}
