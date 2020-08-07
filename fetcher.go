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

type response struct {
	Body   []byte
	Status int
}

// Fetch get response from URL(or API) and return it as JSON.
func (r *response) fetch(url ...interface{}) (*response, error) {
	resp, err := http.Get(fmt.Sprint(url...))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r.Status = resp.StatusCode
	respBody, _ := ioutil.ReadAll(resp.Body)
	r.Body = respBody
	return r, nil
}

// Fetch Initialize fetcher
func Fetch(u string) (*response, error) {
	r := response{}
	return r.fetch(u)
}
