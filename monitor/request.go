package monitor

import (
	"net/http"
	"time"
)

// Request : Sends a http request to given url
// if responce comes before crawltimeout then it sends status of responce as string Through channel
// if responce does not come before crawltimeout then it sends "Timeout" Through  Channel
// Crawl time out is in seconds
func Request(url string, crawlTimeout int, channel chan string) {
	timer := time.NewTimer(time.Duration(crawlTimeout) * time.Second)
	requestChannel := make(chan string)
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			panic(err.Error())
		}
		requestChannel <- resp.Status
	}()
	var msg string
	select {
	case <-timer.C:
		msg = "Timeout"
	case msg = <-requestChannel:
	}
	channel <- msg
}
