package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello, this is a playground file")
	const URL string = "https://www.example.com/"

	thisServer := ServerData{
		URL: URL,
	}
	fmt.Println("Reading from: ", thisServer.URL)
	data, res := thisServer.GetRequest()

	defer res.Body.Close()
	fmt.Println("Server data:")
	fmt.Print(data)
}

// get request from an http server
type ServerData struct {
	URL string
}

func (server ServerData) GetRequest() (string, *http.Response) {
	response, err := http.Get(server.URL)
	server.throwError(err)

	// read the response
	databytes, err := io.ReadAll(response.Body)
	server.throwError(err)

	return string(databytes), response
}

func (server ServerData) throwError(err error) {
	if err != nil {
		panic(err)
	}
}
