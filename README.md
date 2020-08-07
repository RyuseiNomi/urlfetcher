# urlFetcher

## About this repository

This repository is a simple package to get API response.

## Usage

You need to run the command as below to get urlfetcher package.

```
$ go get github.com/RyuseiNomi/urlfetcher
```

In your code, import the package and use `Fetch()` method like this.

```go:sample.go
package main

import (
	"log"

	"github.com/RyuseiNomi/urlfetcher"
)

func main() {
	fetcher, err := urlfetcher.Fetch(apiSampleUrl)
	if err != nil {
		log.Fatal("Unexpected Error : %s", err)
	}
	if fetcher.Status != 200 {
		log.Fatal("Status Error : Got response %v", fetcher.Status)
	}
	if fetcher.Body == nil {
		log.Fatal("Failed to assert : Response body is nil.")
    }
    
    log.Printf("result %v:", fetcher.Body)
}
```