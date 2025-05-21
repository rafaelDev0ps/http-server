package controller

import (
	"http-server/app/request"
	"http-server/app/response"
)

// Return response object instead (int, error)
func DefaultController(request request.Request) response.Response {
	var res response.Response
	res.ResponseHeader = make(response.ResponseHeader)

	if request.RequestHeaders["Connection"] == "close" {
		res.AddHeader("Connection", "close")
	}

	res.ProtocolVersion = "HTTP/1.1"
	res.StatusCode = "200 OK"

	return res
}
