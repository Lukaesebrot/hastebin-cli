package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Paste pastes the given string to the given hastebin instance and returns the URL of the haste
func Paste(content, instance string) (string, error) {
	// Make a POST request to the instances documents endpoint
	instanceAddress := fmt.Sprintf("%s/%s", instance, "documents")
	httpResponse, err := http.Post(instanceAddress, "text/plain", strings.NewReader(content))
	if err != nil {
		return "", err
	}

	// Read the body of the http response
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return "", err
	}

	// Parse the body into a response struct
	parsedResponse := new(response)
	err = json.Unmarshal(body, &parsedResponse)
	if err != nil {
		return "", err
	}

	// Build the URL and return it
	return fmt.Sprintf("%s/%s", instance, parsedResponse.Key), nil
}
