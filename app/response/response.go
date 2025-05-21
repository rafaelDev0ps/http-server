package response

type Response struct {
	StatusCode string
	Header     map[string]string
	Body       []byte
}

const (
	HTTP200 = "200 OK"
	HTTP201 = "201 Created"
	HTTP404 = "404 Not Found"
	HTTP500 = "500 Internal Server Error"
)

func (res *Response) ParseReponse() []byte {
	resp := []byte("HTTP/1.1 " + res.StatusCode + "\r\n")
	if len(res.Header) > 0 {
		for key, value := range res.Header {
			header := key + ": " + value + "\r\n"
			resp = append(resp, header...)
		}
	}
	resp = append(resp, "\r\n"...) //// CRLF that marks the end of the headers

	if len(res.Body) > 0 {
		resp = append(resp, res.Body...)
	}

	return resp
}

func (res *Response) AddHeader(key string, value string) {
	res.Header[key] = value
}
