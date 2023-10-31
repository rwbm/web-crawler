package downloader

import (
	"log"
	"web-crawler/common"
)

// FecthWebPage reads pages to process from workerChan.
// It returns a channel where the results are returned
// and another channel to send any error that might happen.
func FecthWebPage(client common.HttpClient, maxWorkers int, workerChan chan common.WebPage) (chan common.WebPage, chan error) {
	resultsChan := make(chan common.WebPage, maxWorkers)
	errChan := make(chan error)

	for i := 0; i < maxWorkers; i++ {
		go func() {
			for page := range workerChan {

				log.Printf("getting page from %s", page.Url)

				result, err := client.Get(page.Url)
				if err != nil {
					errChan <- err
					continue
				}
				page.Body = result.Body
				resultsChan <- page
			}
		}()
	}
	return resultsChan, errChan
}
