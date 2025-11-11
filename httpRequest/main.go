package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	HTTPServer := HTTPRequest{
		URL:        "https://www.example.com",
		ThrowError: throwError,
	}

	fmt.Println("Welcome to this HTTP server")
	requestBody, response := HTTPServer.GETRequest()
	defer response.Body.Close()

	fmt.Printf("Below is the Request body:\n---------------\n")
	fmt.Println(requestBody)

}

// server type
type HTTPRequest struct {
	URL        string
	ThrowError func(err error)
}

func (h HTTPRequest) GETRequest() (string, *http.Response) {
	response, err := http.Get(h.URL)
	h.ThrowError(err)

	// parsing the error
	databytes, err := io.ReadAll(response.Body)
	h.ThrowError(err)
	return string(databytes), response
}

func throwError(err error) {
	if err != nil {
		panic(err)
	}
}
