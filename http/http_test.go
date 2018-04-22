package http

import (
	"net/http"
	"fmt"
	"time"
	"testing"
	"runtime"
)

var urls = []string{
	"http://pulsoconf.co/",
	"http://golang.org/",
	"http://matt.aimonetti.net/",
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func asyncHttpGets(urls []string, ch chan *HttpResponse)  {

	for _, url := range urls {
		go func(url string) {
			time.Sleep(5*time.Second)
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			resp.Body.Close()
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}

}


func TestAsynHttp(t *testing.T) {
	ch := make(chan *HttpResponse, len(urls)) // buffered

	asyncHttpGets(urls, ch)

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			fmt.Println("url",r.response.Status)
		case <-time.After(10 * time.Second):
			fmt.Println("cpu", runtime.NumCPU())
			fmt.Print("Timeout over")
			return
		}
	}
}






