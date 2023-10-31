package common

import "net/http"

// HttpPClient provides a simple abstraction of HTTP get operation.
type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}
