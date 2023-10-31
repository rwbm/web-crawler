package crawler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"web-crawler/common"
	"web-crawler/downloader"
	"web-crawler/graph"
	"web-crawler/parser"
)

// NewCrawler returns a new instance of Crawler.
func NewCrawler(maxWorkers int) *Crawler {
	return &Crawler{
		maxWorkers: maxWorkers,
		graph:      graph.NewGraph(),
		client:     http.DefaultClient,
	}
}

// Crawler is the main component that coordinates the crawling process.
type Crawler struct {
	urlWorkerChannel chan common.WebPage
	maxWorkers       int
	graph            *graph.Graph
	originalUrl      string
	baseUrl          *url.URL
	client           common.HttpClient
	pagesToProcess   int
}

// Crawl starts the crawling process.
func (c *Crawler) Crawl(webUrl string) error {

	c.originalUrl = webUrl
	baseUrl, err := url.Parse(webUrl)
	if err != nil {
		return fmt.Errorf("%s: %s", webUrl, err)
	}

	c.baseUrl = baseUrl

	// channel to send urls to be processed
	c.urlWorkerChannel = make(chan common.WebPage)

	// channel to get page data
	fetchChannel, errChan := downloader.FecthWebPage(c.client, c.maxWorkers, c.urlWorkerChannel)

	// handle errors from fetching the pages
	go c.handleErrors(errChan)

	// channel to get urls extracted from pages
	resultChan := parser.GetLinks(fetchChannel)

	// put the first URL in the worker channel
	c.urlWorkerChannel <- common.WebPage{Url: webUrl, Body: nil, Links: nil}
	c.graph.AddNode(webUrl)
	c.pagesToProcess = 1

	// keep working while we have something to process
	for c.pagesToProcess > 0 && len(resultChan) == 0 {

		log.Printf("pages in the queue to process: %v", c.pagesToProcess)

		result := <-resultChan
		c.pagesToProcess--

		for _, link := range result.Links {

			// parse url and validate it
			parsedUrl, err := url.Parse(link)
			if err != nil {
				log.Printf("%s: %s", link, err)
				continue
			}

			if !c.isAbsoluteUrl(parsedUrl) {
				// asume current domain
				parsedUrl.Scheme = c.baseUrl.Scheme
				parsedUrl.Host = c.baseUrl.Host
				base := fmt.Sprintf("%s://%s", c.baseUrl.Scheme, c.baseUrl.Host)
				if link, err = url.JoinPath(base, parsedUrl.Path); err != nil {
					log.Printf("%s: %s", link, err)
					continue
				}
			}

			link = strings.TrimSuffix(link, "/")

			if !c.graph.HasNode(link) {
				// validate domain and send page to process
				if parsedUrl.Host == c.baseUrl.Host {
					c.pagesToProcess++
					c.urlWorkerChannel <- common.WebPage{Url: link, Body: nil, Links: nil}
				}

				c.graph.AddNode(link)
				c.graph.AddEdge(result.Url, link)
				continue
			}

			c.graph.AddEdge(result.Url, link)
		}
	}
	return nil
}

func (c *Crawler) handleErrors(errChan chan error) {
	for err := range errChan {
		c.pagesToProcess--
		fmt.Println(err)
	}
}

func (c *Crawler) PrintSiteMap(output io.Writer) {
	c.graph.Print(output, c.originalUrl)
}

func (c *Crawler) isAbsoluteUrl(u *url.URL) bool {
	if u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}
