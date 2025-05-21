package controller

import (
	"fmt"
	"log/slog"
	"strings"

	"http-server/app/request"
	"http-server/app/response"
	"http-server/app/utils"
)

func EchoController(request request.Request) response.Response {
	var res response.Response
	res.ResponseHeader = make(response.ResponseHeader)

	if request.RequestHeaders["Connection"] == "close" {
		res.AddHeader("Connection", "close")
	}

	arg := strings.TrimPrefix(request.Path, "/echo/")

	if arg == "" || arg == request.Path {
		res.StatusCode = "404 Not Found"
		res.ProtocolVersion = "HTTP/1.1"
		return res
	}

	if strings.Contains(request.RequestHeaders["Accept-Encoding"], "gzip") {
		compressedBody, err := utils.CompressContent(arg)
		if err != nil {
			slog.Error(err.Error())
			res.StatusCode = "500 Internal Server Error"
			res.ProtocolVersion = "HTTP/1.1"
			return res
		}

		res.StatusCode = "200 OK"
		res.ProtocolVersion = "HTTP/1.1"
		res.AddHeader("Content-Type", "text/plain")
		res.AddHeader("Content-Encoding", "gzip")
		res.AddHeader("Content-Length", fmt.Sprint(len(compressedBody)))
		res.Body = compressedBody
		return res
	}

	res.StatusCode = "200 OK"
	res.ProtocolVersion = "HTTP/1.1"
	res.AddHeader("Content-Type", "text/plain")
	res.AddHeader("Content-Length", fmt.Sprint(len(arg)))
	res.Body = []byte(arg)

	return res
}
