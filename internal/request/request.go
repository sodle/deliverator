package request

import (
	"fmt"
	"net/http"
)

type Request struct {
  Name    string;
  Method  string;
  Url     string;
}

func (r *Request) Send() (*http.Response, error) {
  request, err := http.NewRequest(r.Method, r.Url, http.NoBody)
  if err != nil {
    return nil, fmt.Errorf("Failed to set up request %w", err)
  }

  return http.DefaultClient.Do(request)
}

