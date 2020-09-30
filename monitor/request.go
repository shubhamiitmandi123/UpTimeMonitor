package monitor

import (
	"fmt"
	"net/http"
	"time"
)

// Request : Sends a http request to given url
// if responce comes before crawltimeout then it sends status of responce as string Through channel
// if responce does not come before crawltimeout then it sends "Timeout" Through  Channel
// Crawl time out is in seconds
func Request(url string, crawlTimeout int, channel chan string) {
	timer := time.NewTimer(time.Duration(crawlTimeout) * time.Second) // timer is set to crawl Timeout
	//Routine which sends http request will send status on requestChannel
	requestChannel := make(chan string)
	go func() {
		resp, err := http.Get(url) // send http request
		if err != nil {
			fmt.Println("ERROR! Request to: ", url, " failed")
		} else {
			requestChannel <- resp.Status //send status to waiting routine
		}
	}()
	var msg string
	select {
	case <-timer.C: //if Crawle timeout
		msg = "Timeout"
	case msg = <-requestChannel: //if respoce comes
	}
	channel <- msg //send Result of request to monitoring routine
}
