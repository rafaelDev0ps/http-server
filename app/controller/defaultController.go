package controller

import (
	"http-server/app/request"
	"http-server/app/response"
)

// Return response object instead (int, error)
func DefaultController(request request.Request) response.Response {
	var res response.Response
	res.Header = make(map[string]string)

	if request.Header["Connection"] == "close" {
		res.AddHeader("Connection", "close")
	}
	res.StatusCode = "200 OK"

	return res
}
