package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web-crawler/crawler"
)

const (
	minWorkers = 2
)

var (
	website = flag.String("url", "", "URL of the website to crawl")
	workers = flag.Int("workers", minWorkers, "Max number of workers to run in parallel")
)

func main() {

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		fmt.Println("process terminated.")
		os.Exit(0)
	}()

	// get website urls
	flag.Parse()
	siteURL := getWebsiteParameter()
	if siteURL == "" {
		printUsage("url")
		os.Exit(1)
	}

	// validate workers
	if workers != nil {
		if *workers < minWorkers {
			*workers = minWorkers
		}
	} else {
		*workers = minWorkers
	}

	start := time.Now()
	log.Printf("starting to crawl %s", siteURL)
	crawler := crawler.NewCrawler(*workers)
	crawler.Crawl(siteURL)
	log.Printf("process completed in %s. Printing results", time.Since(start))
	time.Sleep(1 * time.Second)

	// print result
	crawler.PrintSiteMap(os.Stdout)
}

func getWebsiteParameter() string {
	if website != nil && len(*website) > 0 {
		return *website
	}
	return ""
}

func printUsage(param string) {
	fmt.Printf("missing parameter '%s'\n", param)
	flag.Usage()
}
