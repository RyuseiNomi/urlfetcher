package example

import (
	"log"

	"github.com/RyuseiNomi/urlfetcher"
)

var apiSampleURL = "https://zipcloud.ibsnet.co.jp/api/search?zipcode=1000000"

func getAPIResponseByBytes() {
	resp, err := urlfetcher.Fetch(apiSampleURL)
	if err != nil {
		log.Fatal(err)
	}
	if resp.Status != 200 {
		log.Fatal("Response status is not 200.")
	}
	log.Printf("Response as bytes: %s", resp.Body)
}
