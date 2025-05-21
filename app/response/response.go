package response

type Response struct {
	StatusCode string
	Header     map[string]string
	Body       []byte
}

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
