package urlfetcher

import (
	"testing"
)

var (
	apiSampleUrl = "https://zipcloud.ibsnet.co.jp/api/search?zipcode=1000000"
)

func TestFetch(t *testing.T) {
	fetcher, err := Fetch(apiSampleUrl)
	if err != nil {
		t.Errorf("Unexpected Error : %s", err)
	}
	if fetcher.status != 200 {
		t.Errorf("Status Error : Got response %v", fetcher.status)
	}
	if fetcher.body == nil {
		t.Errorf("Failed to assert : Response body is nil.")
	}
}
