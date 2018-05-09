package services

import (
	"time"
	"encoding/json"
	"net/http"
	"errors"
	"io/ioutil"
	"bytes"

	"github.com/writetosalman/go-rest-api-boilerplate/utilities"
)

var jsonClient 	= &http.Client{Timeout: 20 * time.Second}

func JsonGet(pointerToTarget interface{}, url string) error {

	// Get URL
	r, err := jsonClient.Get(url)
	if err != nil {
		return errors.New("Error occurred while opening URL. "+err.Error())
	}

	// Do not forget to close connection
	defer r.Body.Close()

	// Decode the response and return
	return json.NewDecoder(r.Body).Decode(pointerToTarget)
}

func JsonPost(pointerToTarget interface{}, url string, postData interface{}) error {
	// Marshall post data into JSON
	resultBytes, _ 	:= json.Marshal(postData)

	// POST data to http
	response, err 	:= jsonClient.Post(url, "application/json", bytes.NewBuffer(resultBytes))

	// Check if error occurred
	if err != nil {
		utilities.Log("JsonPost: HTTP post failed with error" + err.Error())
		return errors.New("HTTP post failed with error" + err.Error())
	}

	if response.StatusCode != 200 {
		utilities.Log("JsonPost: HTTP post was not successful [Status: " + utilities.IntToString(response.StatusCode) + "].")
		return errors.New("HTTP post was not successful [Status: " + utilities.IntToString(response.StatusCode) + "]")
	}

	// Read response
	data, _ := ioutil.ReadAll(response.Body)

	// Unmarshall response
	return json.Unmarshal(data, &pointerToTarget)

/*
	// This the implementation web3-go project //

	// Parse Post DATA into JSON
	resultBytes, err := json.Marshal(postData)

	// Check for error in parsing
	if err != nil {
		utilities.Log("JsonPost: Failed to parse HTTP post data. " + err.Error())
		return errors.New("Failed to parse HTTP post data. "+ err.Error())
	}

	// Post Request
	utilities.Log("--------------------------")
	utilities.Log("POST Data")
	utilities.Log(string(resultBytes))
	utilities.Log("--------------------------")

	// Build body and post
	body 		:= strings.NewReader(string(resultBytes))

	// Create New Request
	req, err 	:= http.NewRequest("POST", url, body)

	// Check If error occurred in creating new request
	if err != nil {
		utilities.Log("JsonPost: Failed to create HTTP request. " + err.Error())
		return errors.New("Failed to create HTTP request. "+ err.Error())
	}

	// Set Headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	// Create HTTP client
	var netClient = &http.Client{
		Timeout: time.Second * time.Duration(10),
	}

	// Perform Request over HTTP
	response, err := netClient.Do(req)

	// Check if ERROR occurred
	if err != nil {
		utilities.Log("JsonPost: Failed to send request via HTTP. " + err.Error())
		return errors.New("Failed to send request via HTTP. "+ err.Error())
	}

	// Defer Close
	defer response.Body.Close()

	var bodyBytes []byte

	// Check if HTTP status is not OK
	if response.StatusCode != 200 {
		utilities.Log("JsonPost: HTTP post was not successful [Status: " + utilities.IntToString(response.StatusCode) + "].")
		return errors.New("HTTP post was not successful [Status: " + utilities.IntToString(response.StatusCode) + "]")
	}

	// Read Response
	bodyBytes, err = ioutil.ReadAll(response.Body)

	// Error occurred while reading response
	if err != nil {
		utilities.Log("JsonPost: error occurred while reading response. " + err.Error())
		return errors.New("JsonPost: error occurred while reading response. " + err.Error())
	}

	// Return Unmarshal data
	return json.Unmarshal(bodyBytes, pointerToTarget)
*/
}