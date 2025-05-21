package controller

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"strings"

	"http-server/app/request"
	"http-server/app/response"
	"http-server/app/utils"
)

func FilesController(request request.Request) response.Response {
	var res response.Response
	res.ResponseHeader = make(response.ResponseHeader)

	if request.RequestHeaders["Connection"] == "close" {
		res.AddHeader("Connection", "close")
	}

	filename := strings.TrimPrefix(request.Path, "/files/")
	fileDir := "/tmp/"

	if request.Method == "POST" {
		err := utils.WriteFile(fileDir+filename, string(request.RequestBody))
		if err != nil {
			slog.Error(err.Error())
			res.StatusCode = "500 Internal Server Error"
			res.ProtocolVersion = "HTTP/1.1"
			return res

		}
		res.StatusCode = "201 Created"
		res.ProtocolVersion = "HTTP/1.1"
		return res
	}

	fileContent, err := utils.ReadFile(fileDir + filename)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		res.StatusCode = "404 Not Found"
		res.ProtocolVersion = "HTTP/1.1"
		return res
	}

	res.StatusCode = "200 OK"
	res.ProtocolVersion = "HTTP/1.1"
	res.AddHeader("Content-Type", "application/octet-stream")
	res.AddHeader("Content-Length", fmt.Sprint(len(fileContent)))
	res.Body = fileContent

	return res
}
