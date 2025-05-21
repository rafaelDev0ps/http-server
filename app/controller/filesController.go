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
	res.Header = make(map[string]string)

	if request.Header["Connection"] == "close" {
		res.AddHeader("Connection", "close")
	}

	filename := strings.TrimPrefix(request.Path, "/files/")
	fileDir := "/tmp/"

	if request.Method == "POST" {
		err := utils.WriteFile(fileDir+filename, string(request.Body))
		if err != nil {
			slog.Error(err.Error())
			res.StatusCode = response.HTTP500
			return res

		}
		res.StatusCode = response.HTTP201
		return res
	}

	fileContent, err := utils.ReadFile(fileDir + filename)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		res.StatusCode = response.HTTP404
		return res
	}

	res.StatusCode = response.HTTP200
	res.AddHeader("Content-Type", "application/octet-stream")
	res.AddHeader("Content-Length", fmt.Sprint(len(fileContent)))
	res.Body = fileContent

	return res
}
