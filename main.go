/*
package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i< len(people); i++ {
		fmt.Println("Waiting for", i)
		fmt.Println("Received this message:", <-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}


func Count(person string) {
	for i:=0;i<10;i++{
		fmt.Println(person, "is good", i)
		time.Sleep(time.Second)
	}
}
*/
package main

import (
	"errors"
	"fmt"
	"net/http"
)

type request struct {
	url string
	status string
}

var errRequestFailed = errors.New("Request Failed")

func main() {
	results := make(map[string]string)
	c := make(chan request)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i:=0; i<len(urls); i++{
		result :=<-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println("URL: " + url + " STATUS: " + status)
	}
}


func hitURL(url string, c chan<- request) { // send only
	//fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	} 
	c <- request{url: url, status: status}
}