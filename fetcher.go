package urlfetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Fetcher The package for get API response
type Fetcher interface {
	Fetch(...interface{})
}

type fetcher struct {
	Body   []byte
	Status int
}

// Fetch get response from URL(or API) and return it as JSON.
func (f *fetcher) fetch(url ...interface{}) (*fetcher, error) {
	resp, err := http.Get(fmt.Sprint(url...))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	f.Status = resp.StatusCode
	respBody, _ := ioutil.ReadAll(resp.Body)
	f.Body = respBody
	return f, nil
}

// Fetch Initialize fetcher
func Fetch(u string) (*fetcher, error) {
	f := fetcher{}
	return f.fetch(u)
}
