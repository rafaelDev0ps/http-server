package controller

import (
	"fmt"
	"net"

	"http-server/app/request"
	"http-server/app/response"
)

func UserAgentController(conn net.Conn, request request.Request) (int, error) {
	var res response.Response
	res.ResponseHeader = make(response.ResponseHeader)

	if request.RequestHeaders["Connection"] == "close" {
		res.AddHeader("Connection", "close")
	}

	res.AddHeader("Content-Type", "text/plain")
	res.AddHeader("Content-Length", fmt.Sprint(len(request.RequestHeaders["User-Agent"])))

	return fmt.Fprintf(conn, "%s%s\r\n%s", res.Status200(), res.FormatHeaders(), request.RequestHeaders["User-Agent"])
}
