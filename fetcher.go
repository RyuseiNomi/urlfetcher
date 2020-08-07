package urlfetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Fetcher API取得に関する動作のパッケージ
type Fetcher interface {
	Fetch(...interface{})
}

type fetcher struct {
	body   []byte
	status int
}

// Fetch get response from URL(or API) and return it as JSON.
func (f *fetcher) fetch(url ...interface{}) (*fetcher, error) {
	resp, err := http.Get(fmt.Sprint(url...))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	f.status = resp.StatusCode
	respBody, _ := ioutil.ReadAll(resp.Body)
	f.body = respBody
	return f, nil
}

// Fetch Initialize fetcher
func Fetch(u string) (*fetcher, error) {
	f := fetcher{}
	return f.fetch(u)
}
