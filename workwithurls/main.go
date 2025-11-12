package main

import (
	"fmt"
	"net/url"
)

const TEST_URL = "https://developer.hashicorp.com/tutorials/library?product=terraform"

func main() {
	fmt.Println("Working with URLS")
	fmt.Println("Url: ", TEST_URL)

	// parsing url
	parsedUrl, _ := url.Parse(TEST_URL)
	fmt.Println("Parsed format: ", parsedUrl)

	// url data
	NewUrl := URLDetails{
		Scheme:      parsedUrl.Scheme,
		Host:        parsedUrl.Host,
		Port:        parsedUrl.Port(),
		QueryParams: parsedUrl.Query(),
	}

	// printing the details values
	fmt.Println(NewUrl.Scheme, NewUrl.Host, NewUrl.Port)
	for key, val := range NewUrl.QueryParams {
		fmt.Println(key, ": ", val)
	}

	// reconstructring url
	reUrl := &url.URL{
		Scheme:  NewUrl.Scheme,
		Host:    NewUrl.Host,
		Path:    parsedUrl.Path,
		RawPath: "user=mark",
	}
	fmt.Println("Reconstructed url: ", reUrl.String())
}

type URLDetails struct {
	Scheme      string
	Host        string
	Port        string
	QueryParams url.Values
}
