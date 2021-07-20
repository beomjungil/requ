package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-requ/requ/model"
)

func Request(requestConfig model.HttpRequestConfig) {
	client := &http.Client{}

	request, err := http.NewRequest(
		requestConfig.Method,
		requestConfig.Url,
		bytes.NewBuffer([]byte(requestConfig.Body)),
	)
	if err != nil {
		return
	}

	for key, value := range requestConfig.Headers {
		request.Header.Add(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return
	}

	defer response.Body.Close()

	printResponse(response)
}

func printResponse(response *http.Response) {
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var jsonBody map[string]interface{}
	err = json.Unmarshal(respBody, &jsonBody)
	if err != nil {
		return
	}

	prettyBody, err := json.MarshalIndent(jsonBody, "", "\t")
	if err != nil {
		return
	}

	fmt.Println("Status:" + "\t" + response.Status)
	fmt.Println("Body:")
	fmt.Println(string(prettyBody))
}
