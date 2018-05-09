package controllers

import (
	"net/http"

	"github.com/writetosalman/go-rest-api-boilerplate/utilities"
)

/**
  * Homepage is a function that handles route
  * @param 	http.ResponseWriter
  * @param 	*http.Request
  */
func Homepage(w http.ResponseWriter, r *http.Request) {
	utilities.JsonResponse("Error: 404 - page not found", w, http.StatusNotAcceptable)
	utilities.Log("Error: 404 - page not found")
}

/**
  * Ping is a function that handles route to return a ping request
  * @param 	http.ResponseWriter
  * @param 	*http.Request
  */
func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()

		message 	:= r.Form.Get("message")

		if message == "bad" {
			utilities.JsonResponse("Ping: BAD", w, http.StatusInternalServerError)
		}

		utilities.Log("Ping: - "+message)
		utilities.JsonResponse("Ping - "+message, w, http.StatusOK)

		return
	}

	utilities.Log("Ping: Method not supported")
	utilities.JsonResponse("Ping", w, http.StatusNotAcceptable)
}
