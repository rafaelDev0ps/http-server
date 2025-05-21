package controller

import (
	"fmt"

	"http-server/app/request"
	"http-server/app/response"
)

func UserAgentController(request request.Request) response.Response {
	var res response.Response
	res.ResponseHeader = make(response.ResponseHeader)

	if request.RequestHeaders["Connection"] == "close" {
		res.AddHeader("Connection", "close")
	}

	res.AddHeader("Content-Type", "text/plain")
	res.AddHeader("Content-Length", fmt.Sprint(len(request.RequestHeaders["User-Agent"])))
	res.StatusCode = "200 OK"
	res.ProtocolVersion = "HTTP/1.1"
	res.Body = []byte(request.RequestHeaders["User-Agent"])

	return res
}
