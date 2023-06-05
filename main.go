package main

import (
	"fmt"
	"net/http"
)

type Pinger struct {
	client *http.Client
}

func (p Pinger) Ping(url string) bool {
	resp, err := p.client.Head(url)
	if err != nil {
		return false
	}

	if resp.StatusCode != 200 {

		return false

	}

	return true

}

func main() {
	fmt.Println("Test")

	client := &http.Client{}

	pinger := Pinger{client}

	got := pinger.Ping("https://example.com")

	fmt.Println(got)

}
