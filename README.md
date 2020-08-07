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

import github.com/RyuseiNomi/urlfetcher

func fetchAPI() {
	fetcher, err := urlfetcher.Fetch(apiSampleUrl)
	if err != nil {
		t.Errorf("Unexpected Error : %s", err)
	}
	if fetcher.status != 200 {
		t.Errorf("Status Error : Got response %v", fetcher.status)
	}
	if fetcher.body == nil {
		t.Errorf("Failed to assert : Response body is nil.")
    }
    
    fmt.Printf("result %v:", fetcher.body)
}
```