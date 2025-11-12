package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const LOCAL_URL = "http://127.0.0.1:8000"

func main() {
	fmt.Println("GET Request on an http server")
	message, code := GetRootMessage(LOCAL_URL)
	fmt.Println("Status code: ", code)
	fmt.Println("Message: ", message)
}

func GetRootMessage(url string) (string, int) {
	response, err := http.Get(url)
	handleError(err) // if there's and error

	// get the response
	data, _ := io.ReadAll(response.Body)
	var responseData strings.Builder
	databytes, _ := responseData.Write(data)
	fmt.Println("Databytes count: ", databytes)
	return responseData.String(), response.StatusCode
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
