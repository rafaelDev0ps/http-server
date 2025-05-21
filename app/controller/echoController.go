package controller

import (
	"fmt"
	"log/slog"
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

	request := req.ParseRequest(content)

	if request.RequestHeaders["Connection"] == "close" {
		res.AddHeader("Connection", "close")
	}

	arg := strings.TrimPrefix(request.Path, "/echo/")

	if arg == "" || arg == request.Path {
		return res.Status404(conn)
	}

	if strings.Contains(request.RequestHeaders["Accept-Encoding"], "gzip") {
		compressedBody, err := utils.CompressContent(arg)
		if err != nil {
			slog.Error(err.Error())
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
