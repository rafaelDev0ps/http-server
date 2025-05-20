package controller

import (
	"errors"
	"fmt"
	"io/fs"
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

	path := req.GetPath(content)
	method := req.GetMethod(content)
	body := req.GetBody(content)

	headers := req.GetHeaders(content)

	if req.GetHeaderValue(headers, "Connection") == "close" {
		res.AddHeader("Connection", "close")
	}

	filename := strings.TrimPrefix(path, "/files/")
	fileDir := "/tmp/"

	if method == "POST" {
		err := utils.WriteFile(fileDir+filename, body)
		if err != nil {
			fmt.Println(err)
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
