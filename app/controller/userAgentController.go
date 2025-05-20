package controller

import (
	"fmt"
	"net"
	"strings"

	"http-server/app/request"
	"http-server/app/response"
)

func UserAgentController(conn net.Conn, content []string) (int, error) {
	var req request.Request
	var res response.Response
	res.ResponseHeader = make(response.ResponseHeader)

	headers := req.GetHeaders(content)

	if req.GetHeaderValue(headers, "Connection") == "close" {
		res.AddHeader("Connection", "close")
	}

	var userAgent string
	for i := range len(headers) {
		if strings.Contains(headers[i], "User-Agent") {
			userAgent = headers[i]

		}
	}
	headerUserAgent := strings.Split(userAgent, " ")[1]

	res.AddHeader("Content-Type", "text/plain")
	res.AddHeader("Content-Length", fmt.Sprint(len(headerUserAgent)))

	return fmt.Fprintf(conn, "%s%s\r\n%s", res.Status200(), res.FormatHeaders(), headerUserAgent)
}
