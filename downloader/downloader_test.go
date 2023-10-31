package downloader

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"web-crawler/common"
)

type HttpClientMock struct {
	Response *http.Response
	Error    error
}

func (h *HttpClientMock) Get(url string) (resp *http.Response, err error) {
	if h.Error != nil {
		return nil, h.Error
	}
	return h.Response, nil
}

func TestFetchWithError(t *testing.T) {
	err := errors.New("Mock error")
	client := HttpClientMock{nil, err}
	concurrency := 1
	workChan := make(chan common.WebPage)

	_, errChan := FecthWebPage(&client, concurrency, workChan)

	page := common.WebPage{Url: "https://mywebsite.com"}

	go func() { workChan <- page }()

	outErr := <-errChan

	if outErr == nil {
		t.Error("FecthWebPage failed to return an error")
		return
	}

	if outErr.Error() != err.Error() {
		t.Error("FecthWebPage returned an invalid error")
		return
	}
}

func TestFetch(t *testing.T) {
	body := strings.NewReader(`<html><body/></html>`)
	response := http.Response{Body: io.NopCloser(body)}
	client := HttpClientMock{&response, nil}
	concurrency := 1
	workChan := make(chan common.WebPage)

	outChan, _ := FecthWebPage(&client, concurrency, workChan)

	page := common.WebPage{Url: "https://mywebsite.com"}

	go func() { workChan <- page }()

	outPage := <-outChan

	if outPage.Url != page.Url {
		t.Error("FecthWebPage changed page URL")
		return
	}

	if outPage.Body == nil {
		t.Error("FecthWebPage failed to set body")
		return
	}
}
