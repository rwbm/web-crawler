package crawler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCrawl(t *testing.T) {
	mockedServer := mockHTTPServer()

	c := NewCrawler(1)
	c.Crawl(mockedServer.URL)

	c.PrintSiteMap(os.Stdout)
}

func mockHTTPServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			html   string
			status int
		)

		switch r.URL.Path {

		case "", "/":
			html = `
			<html>
			<body>
				<h2>Home</h2>
				<ul>
				  <li><a href="/foo">Foo</a></li>
				  <li><a href="/bar">Bar</a></li>
				  <li><a href="/zzz">Zzz</a></li>
				</ul>
			</body>
			</html>
			`
			status = http.StatusOK

		case "/foo":
			html = `
			<html>
			<body>
				<h2>Foo</h2>
				<ul>
				  <li><a href="/">Home</a></li>
				  <li><a href="/bar">Bar</a></li>
				</ul>
			</body>
			</html>
			`
			status = http.StatusOK

		case "/bar":
			html = `
			<html>
			<body>
				<h2>Bar</h2>
				<ul>
				  <li><a href="/">Home</a></li>
				  <li><a href="/zzz">Bar</a></li>
				</ul>
			</body>
			</html>
			`
			status = http.StatusOK

		case "/zzz":
			html = `
			<html>
			<body>
				<h2>Zzz</h2>
				<ul>
				  <li><a href="/">Home</a></li>
				  <li><a href="http://google.com">Bar</a></li>
				</ul>
			</body>
			</html>
			`
			status = http.StatusOK

		default:
			status = http.StatusNotFound

		}

		w.WriteHeader(status)

		if html != "" {
			fmt.Fprint(w, html)
		}
	}))

	return server
}
