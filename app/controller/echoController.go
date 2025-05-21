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
	res.Header = make(map[string]string)

	if request.Header["Connection"] == "close" {
		res.AddHeader("Connection", "close")
	}

	arg := strings.TrimPrefix(request.Path, "/echo/")

	if arg == "" || arg == request.Path {
		res.StatusCode = response.HTTP404
		return res
	}

	if strings.Contains(request.Header["Accept-Encoding"], "gzip") {
		compressedBody, err := utils.CompressContent(arg)
		if err != nil {
			slog.Error("error compressing content", "error", err.Error())
			res.StatusCode = response.HTTP500
			return res
		}

		res.StatusCode = response.HTTP200
		res.AddHeader("Content-Type", "text/plain")
		res.AddHeader("Content-Encoding", "gzip")
		res.AddHeader("Content-Length", fmt.Sprint(len(compressedBody)))
		res.Body = compressedBody
		return res
	}

	res.StatusCode = response.HTTP200
	res.AddHeader("Content-Type", "text/plain")
	res.AddHeader("Content-Length", fmt.Sprint(len(arg)))
	res.Body = []byte(arg)

	return res
}
