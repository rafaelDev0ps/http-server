package controller

import (
	"fmt"

	"http-server/app/request"
	"http-server/app/response"
)

func UserAgentController(request request.Request) response.Response {
	var res response.Response
	res.Header = make(map[string]string)

	if request.Header["Connection"] == "close" {
		res.AddHeader("Connection", "close")
	}

	res.AddHeader("Content-Type", "text/plain")
	res.AddHeader("Content-Length", fmt.Sprint(len(request.Header["User-Agent"])))
	res.StatusCode = "200 OK"
	res.Body = []byte(request.Header["User-Agent"])

	return res
}
