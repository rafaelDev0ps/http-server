package response

type Response struct {
	StatusCode      string
	ProtocolVersion string
	ResponseHeader  ResponseHeader
	Body            []byte
}

type ResponseHeader map[string]string

func (res *Response) ParseReponse() []byte {
	resp := []byte(res.ProtocolVersion + " " + res.StatusCode + "\r\n")
	if len(res.ResponseHeader) > 0 {
		for key, value := range res.ResponseHeader {
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
	res.ResponseHeader[key] = value
}
