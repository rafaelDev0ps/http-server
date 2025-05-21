package controller

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net"
	"strings"

	"http-server/app/request"
	"http-server/app/response"
	"http-server/app/utils"
)

func FilesController(conn net.Conn, content []string) (int, error) {
	var req request.Request
	var res response.Response
	res.ResponseHeader = make(response.ResponseHeader)
	request := req.ParseRequest(content)

	if request.RequestHeaders["Connection"] == "close" {
		res.AddHeader("Connection", "close")
	}

	filename := strings.TrimPrefix(request.Path, "/files/")
	fileDir := "/tmp/"

	if request.Method == "POST" {
		err := utils.WriteFile(fileDir+filename, string(request.RequestBody))
		if err != nil {
			slog.Error(err.Error())
			return res.Status500(conn)

		}
		res := &response.Response{StatusCode: "201 Created", ProtocolVersion: "HTTP/1.1"}
		return fmt.Fprintf(conn, "%s\r\n", res.GetResponseLine())

	}

	fileContent, err := utils.ReadFile(fileDir + filename)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		return res.Status404(conn)
	}

	res.AddHeader("Content-Type", "application/octet-stream")
	res.AddHeader("Content-Length", fmt.Sprint(len(fileContent)))
	return fmt.Fprintf(conn, "%s%s\r\n%s", res.Status200(), res.FormatHeaders(), fileContent)
}
