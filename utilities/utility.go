package utilities

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
	"strconv"
	"io/ioutil"
	"errors"
	"os"
)

type Response struct {
	Data string `json:"data"`
}

func SendResponseHeaders(w http.ResponseWriter, httpStatus int) {
	w.Header().Set("Access-Control-Allow-Origin", 	"*")
	w.Header().Set("Access-Control-Allow-Headers", 	"*")
	w.Header().Set("Access-Control-Allow-Methods", 	"GET,POST,OPTIONS")
	w.WriteHeader(httpStatus)
}

// JsonResponse sends out JSON to the browser
func JsonResponse(response string, w http.ResponseWriter, httpStatus int) {
	SendResponseHeaders(w, httpStatus)
	JsonResponseArray(Response{response}, w, httpStatus)
}

// Json Response sents out JSON array to the browser
func JsonResponseArray(response interface{}, w http.ResponseWriter, httpStatus int) {

	jsonParsed, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	SendResponseHeaders(w, httpStatus)
	w.Write(jsonParsed)
}

// JsonResponse sends out JSON to the browser
func JsonResponseParsed(response interface{}, w http.ResponseWriter, httpStatus int) {

	w.Header().Set("Content-Type", "application/json")
	SendResponseHeaders(w, httpStatus)
	json.NewEncoder(w).Encode(response)
}

// Log function writes log messages to the print
func Log(str ...interface{}) {
	fmt.Println(str)
}

// CheckError function checks if there is error or not
func CheckError(err error) {
	if ( err != nil ) {
		Log("Error: "+ err.Error())
		log.Fatal("Error: "+ err.Error())
	}
}

// StringArrayToInterface converts string array to interface array
func StringArrayToInterface(arrString []string) ([]interface{}) {

	// Convert String array To Interface Array | https://golang.org/doc/faq#convert_slice_of_interface
	arrInterface := make([]interface{}, len(arrString))
	for index, item := range arrString {
		arrInterface[index] = item
	}

	return arrInterface
}

// Convert string to integer
func StringToInt(num string) int {

	if num == "" {
		return 0
	}

	number, err := strconv.Atoi(num)
	if err != nil {
		Log("StringToInt Failed. " + err.Error())
		return 0
	}

	return number
}


// Convert string to integer64
func StringToInt64(num string) int64 {

	if num == "" {
		return 0
	}

	number, err := strconv.ParseInt(num, 10, 0)
	if err != nil {
		Log("StringToInt64 Failed. " + err.Error())
		return 0
	}

	return number
}

// Convert string to float64
func StringToFloat64(num string) float64 {

	if num == "" {
		return float64(0)
	}

	number, err := strconv.ParseFloat(num, 0)
	if err != nil {
		Log("StringToFloat64 Failed. " + err.Error())
		return 0
	}

	return number
}

// Convert integer to string
func IntToString(num int) string {

	if num == 0 {
		return "0"
	}

	strNumber := strconv.FormatInt(int64(num), 10)
	return strNumber
}

// Convert integer64 to string
func Int64ToString(num int64) string {

	if num == 0 {
		return "0"
	}

	strNumber := strconv.FormatInt(num, 10)
	return strNumber
}

// Convert float to string
func FloatToString(num float64) string {

	if num == 0 {
		return "0"
	}

	strNumber := strconv.FormatFloat(num, 'f', 10, 64)
	return strNumber
}

// Convert integer to string | https://stackoverflow.com/questions/27137521/how-to-convert-interface-to-string
func InterfaceToString(str interface{}) string {

	if str == nil {
		return ""
	}

	return fmt.Sprintf("%v", str)

	/*
	// Assert if interface{} is string | https://stackoverflow.com/questions/14289256/cannot-convert-data-type-interface-to-type-string-need-type-assertion
	strEmail, isOk 	:= arrTokenClaims["email"].(string)
	if !isOk {
		return map[string]string{"error":"Error parsing token email"}, false
	}*/
}

//func InterfaceToStringMap(arr map[string]interface{}) map[string]string {
//
//	arrString := make(map[string]string, len(arr))
//	for index, item := range arr {
//		arrString[index] = item
//		fmt.Println("Index: " +index)
//		fmt.Println("Item: " +item)
//	}
//	return arrString
//}

// This function writes contents of file
func FileWrite(filename string, data []byte) error {

	if filename == "" {
		Log("FileWrite: Cannot write file. Filename is empty")
		return errors.New("Cannot write file. Filename is empty")
	}

	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		Log("FileWrite: Unable to write file")
		return errors.New("Unable to write file")
	}

	return nil
}

// This function reads contents of file
func FileRead(filename string) ([]byte, error) {

	if filename == ""  {
		Log("FileRead: Cannot read file. Filename is empty")
		return nil, errors.New("Cannot read file. Filename is empty")
	}

	if err := FileExists(filename); err != nil {
		return nil, errors.New("Cannot read file. It does not exists")
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		Log("FileRead: Unable to read file")
		return nil, errors.New("Unable to read file")
	}

	return data, nil
}

// This function checks if file exists
func FileExists(filename string) (error) {

	if filename == ""  {
		Log("FileRead: Cannot check if file exists. Filename is empty")
		return errors.New("Cannot check if file exists. Filename is empty")
	}

	if _, err := os.Stat(filename); err == nil {
		return nil
	}

	return errors.New("File does not exists")
}