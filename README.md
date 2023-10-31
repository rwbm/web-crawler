# Web Crawler

Simple web crawler implementation. Takes a website and extracts all links that belong to the same domain. Results are printed in the console.

## Build and run
```
$ go build -o web-crawler
$ ./web-crawler -url=https://google.com.uy -workers=5
```

## Run tests
```
go test -v -cover ./...
```