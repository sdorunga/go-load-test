package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Verb   string
	URL    string
	client Client
}

type Client interface {
	Do(request Request) (*http.Response, error)
}

type httpClient struct {
	client http.Client
}

func (this *Request) String() string {
	return fmt.Sprintf("%s for %s", this.Verb, this.URL)
}

func (this *httpClient) Do(request Request) (res *http.Response, err error) {
	req, err := http.NewRequest(request.Verb, request.URL, nil)
	if err != nil {
		fmt.Println("Error making a " + request.Verb + " request to " + request.URL)
	}
	res, err = this.client.Do(req)
	return
}

func (this *Request) Perform() {
	res, err := this.client.Do(*this)
	if err != nil {
		fmt.Println("Error making a " + this.Verb + " request to " + this.URL)
		return
	}
	if res.Body != nil {
		defer res.Body.Close()
		ioutil.ReadAll(res.Body)
	}
	return
}
