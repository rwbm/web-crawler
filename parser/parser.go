package parser

import (
	"io"
	"web-crawler/common"

	"golang.org/x/net/html"
)

// GetLinks reads pages from a channel and spins out new go routines
// to handle the tokenisation and extraction of <a /> elementss from html.
func GetLinks(workerChan chan common.WebPage) chan common.WebPage {
	resultsChan := make(chan common.WebPage)
	go func() {
		for page := range workerChan {
			go func(p common.WebPage) {
				if p.Body == nil {
					return
				}
				p.Links = extractLinks(p.Body)
				resultsChan <- p
			}(page)
		}
	}()
	return resultsChan
}

// extractUrls parses all the HTML '<a />' elements and extracts the
// 'href' attribute in each of them.
func extractLinks(body io.Reader) []string {
	foundLinks := make([]string, 0)

	z := html.NewTokenizer(body)
	for {
		tokenType := z.Next()
		if tokenType == html.ErrorToken {
			return foundLinks
		}
		if tokenType == html.StartTagToken || tokenType == html.SelfClosingTagToken {
			token := z.Token()
			if token.Data == "a" {
				for _, a := range token.Attr {
					if a.Key == "href" {
						link := a.Val
						foundLinks = append(foundLinks, link)
					}
				}
			}
		}
	}
}
