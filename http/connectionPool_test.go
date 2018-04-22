package http

import (
	"net/http"
	"testing"
	"fmt"
	"html"
	"io/ioutil"
	"io"
	"time"
)


//var myClient *http.Client
var count = make (chan int,100)
func startWebserver1() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Println("Hello, %q", html.EscapeString(r.URL.Path))
		time.Sleep(300*time.Millisecond)
	})
}

func sendRequest(){
	count <- 1
	fmt.Println("BUffer:",len(count),cap(count))
	resp, err := http.Get("http://localhost:8081/")
	if err != nil {
		panic(fmt.Sprintf("Got error: %v", err))
	}
	//time.Sleep(500*time.Millisecond)
	//discard response
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	//fmt.Println("Response:", resp.Status)
	<- count
}

func startLoadTest1() {
	for {
		time.Sleep(1*time.Millisecond)
		go sendRequest()
	}

}

func TestConnect1(t *testing.T) {

	//tr := &http.Transport{
	//	MaxIdleConns:       100,
	//	IdleConnTimeout:    15 * time.Second,
	//	DisableCompression: true,
	//}
	//myClient = &http.Client{Transport: tr}


	go startLoadTest1()
	startWebserver1()
	err := http.ListenAndServe(":8081", nil)
	if err != nil{
		fmt.Println("ERROR TO START SERVER")
		return
	}
}
