package http

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"testing"
	"io/ioutil"
	"io"
)

func startWebserver() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	go http.ListenAndServe(":8081", nil)

}

func startLoadTest() {
	count := 0
	for {
		resp, err := http.Get("http://localhost:8081/")
		if err != nil {
			panic(fmt.Sprintf("Got error: %v", err))
		}
		//discard response
		io.Copy(ioutil.Discard, resp.Body)

		resp.Body.Close()
		log.Printf("Finished GET request #%v", count)
		count += 1
	}

}

func TestConnect(t *testing.T) {

	// start a webserver in a goroutine
	startWebserver()

	startLoadTest()

}
