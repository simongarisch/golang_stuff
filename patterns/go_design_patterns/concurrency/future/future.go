/*
The Future design pattern (also called Promise) is a quick and easy way to achieve
concurrent structures for asynchronous programming.
See https://github.com/chebyrash/promise
and http://www.golangpatterns.info/concurrency/futures
and https://medium.com/strava-engineering/futures-promises-in-the-land-of-golang-1453f4807945
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RequestFuture(url string) <-chan []byte {
	c := make(chan []byte, 1)

	go func() {
		var body []byte
		defer func() { c <- body }()
		resp, err := http.Get(url)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		body, _ = ioutil.ReadAll(resp.Body)
	}()

	return c
}

func main() {
	future := RequestFuture("https://labs.strava.com/")

	// get the result
	body := <-future
	fmt.Println(body)
}
