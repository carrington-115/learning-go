package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const LOCAL_URL = "http://127.0.0.1:8000"

func main() {
	// Get request
	fmt.Println("GET Request on an http server")
	message, code := GetRootMessage(LOCAL_URL)
	fmt.Println("Status code: ", code)
	fmt.Println("Message: ", message)

	// Post request
	fmt.Println("\nPOST Request on an http server")
	postMsg, pc := SendPostRequest(LOCAL_URL + "/get_package")
	fmt.Println("Status code: ", pc)
	fmt.Println("Message: ", postMsg)

	// post request with form
	fmt.Println("\nPOST request with Encoded form data on an http server")
	postFormRes, pfc := SendPostRequest(LOCAL_URL + "/get_package")
	fmt.Println("Status code: ", pfc)
	fmt.Println("Message: ", postFormRes)
}

func GetRootMessage(serverUrl string) (string, int) {
	response, err := http.Get(serverUrl)
	handleError(err) // if there's and error

	// post the response
	data, _ := io.ReadAll(response.Body)
	var responseData strings.Builder
	databytes, _ := responseData.Write(data)
	fmt.Println("Databytes count: ", databytes)
	return responseData.String(), response.StatusCode
}

func SendPostRequest(serverUrl string) (string, int) {
	bodyData := strings.NewReader(`
			{
			"title": "Test title",
			"weight": 4.5,
			"quantity": 10
			}
	`)

	response, err := http.Post(serverUrl, "application/json", bodyData)
	handleError(err)
	defer response.Body.Close()

	// unique response
	responseData, _ := io.ReadAll(response.Body)
	return string(responseData), response.StatusCode
}

func SendPostRequestWithForm(serverUrl string) (string, int) {
	data := url.Values{}
	data.Add("title", "Test title")
	data.Add("weight", "4.5")
	data.Add("quantity", "10")

	res, err := http.PostForm(serverUrl, data)
	handleError(err)
	resData, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	return string(resData), res.StatusCode
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
