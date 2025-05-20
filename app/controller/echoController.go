package controller

import (
	"fmt"
	"net"
	"strings"

	"http-server/app/request"
	"http-server/app/response"
	"http-server/app/utils"
)

func EchoController(conn net.Conn, content []string) (int, error) {
	var req request.Request
	var res response.Response
	res.ResponseHeader = make(response.ResponseHeader)

	path := req.GetPath(content)

	headers := req.GetHeaders(content)

	if req.GetHeaderValue(headers, "Connection") == "close" {
		res.AddHeader("Connection", "close")
	}

	arg := strings.TrimPrefix(path, "/echo/")

	if arg == "" || arg == path {
		return res.Status404(conn)
	}

	headerAcceptEncoding := req.GetHeaderValue(headers, "Accept-Encoding")

	if strings.Contains(headerAcceptEncoding, "gzip") {
		compressedBody, err := utils.CompressContent(arg)
		if err != nil {
			fmt.Println(err)
			return res.Status500(conn)
		}
		res.AddHeader("Content-Type", "text/plain")
		res.AddHeader("Content-Encoding", "gzip")
		res.AddHeader("Content-Length", fmt.Sprint(len(compressedBody)))
		return fmt.Fprintf(conn, "%s%s\r\n%s", res.Status200(), res.FormatHeaders(), compressedBody)
	}

	res.AddHeader("Content-Type", "text/plain")
	res.AddHeader("Content-Length", fmt.Sprint(len(arg)))
	return fmt.Fprintf(conn, "%s%s\r\n%s", res.Status200(), res.FormatHeaders(), arg)
}
