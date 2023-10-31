package common

import "io"

// WebPage represents a web page to be crawled.
type WebPage struct {
	Url   string
	Body  io.Reader
	Links []string
}
