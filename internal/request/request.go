package request

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Request struct {
	Name   string
	Method string
	Url    string
	Body   string
}

func (r *Request) Send() (*http.Response, error) {
	var body io.Reader
	if r.Body != "" {
		body = strings.NewReader(r.Body)
	} else {
		body = nil
	}

	request, err := http.NewRequest(r.Method, r.Url, body)
	if err != nil {
		return nil, fmt.Errorf("Failed to set up request %w", err)
	}

	return http.DefaultClient.Do(request)
}
