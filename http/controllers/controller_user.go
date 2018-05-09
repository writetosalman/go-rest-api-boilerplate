package controllers

import (
	"net/http"

	AuthenticateService "github.com/writetosalman/go-rest-api-boilerplate/services"
	"github.com/writetosalman/go-rest-api-boilerplate/utilities"
	"github.com/writetosalman/go-rest-api-boilerplate/models"
	"github.com/gorilla/context"
)

// https://www.thepolyglotdeveloper.com/2017/03/authenticate-a-golang-api-with-json-web-tokens/
type JwtToken struct {
	Token string 	`json:"token"`
}

/**
  * Login is a function that handles route /login
  * @param 	http.ResponseWriter
  * @param 	*http.Request
  */
func Login(w http.ResponseWriter, r *http.Request) {
	utilities.Log("Login: Trying")

	if r.Method == "POST" || r.Method == "GET" {
		r.ParseForm()

		password 	:= r.Form.Get("password")
		email 		:= r.Form.Get("email")

		utilities.Log("Login: request", password, email)

		// Check if post data is empty
		if email == "" || password == "" {
			utilities.Log("Login: Data empty")
			utilities.JsonResponse("Login Data empty", w, http.StatusUnauthorized)
			return
		}

		// Try to validate login details
		utilities.Log("Login: email & password provided. Now load User.")
		User, err	:= models.GetUserByEmail(email)
		if err != nil {
			utilities.Log("Login: " + err.Error())
			utilities.JsonResponse("Login: " + err.Error(), w, http.StatusUnauthorized)
			return
		}

		// Verify Password Hash
		if response, isOk := AuthenticateService.VerifyUserPassword(password, User.PasswordHash); !isOk {
			utilities.Log("Login: Incorrect login credentials - " + response)
			utilities.JsonResponse("Incorrect login credentials", w, http.StatusUnauthorized)
			return
		}

		// Credentials are fine, Get Token
		utilities.Log("Login: Credentials verified. UserID: "+ utilities.IntToString(User.UserID))
		tokenString, err := AuthenticateService.GetJwtToken(User.UserID, email)
		if err != nil {
			utilities.Log("Login: Failed to get token - " + err.Error())
			utilities.JsonResponse("Login: Failed to get token - " + err.Error(), w, http.StatusUnauthorized)
			return
		}

		// Return Token
		utilities.Log("Token: " + tokenString)
		utilities.JsonResponseParsed(JwtToken{Token: tokenString}, w, http.StatusOK)
		return
	}
	utilities.Log("Login: Failed")
	utilities.JsonResponse("Login Failed", w, http.StatusNotAcceptable)
}

/**
  * Dashboard is a function that handles route /dashboard
  * @param 	http.ResponseWriter
  * @param 	*http.Request
  */
func Dashboard(w http.ResponseWriter, r *http.Request) {

	// Get information from parsed Token | https://stackoverflow.com/questions/31504456/how-can-i-pass-data-from-middleware-to-handlers
	u := context.Get(r, "User")
	var user models.User
	var isOk bool

	// Type assert | https://stackoverflow.com/questions/48576905/how-to-use-type-struct-saved-in-context-in-middleware-accessed-in-controller/48577245#48577245
	if user, isOk = u.(models.User); !isOk {
		utilities.JsonResponse("Dashboard: User not loaded", w, http.StatusInternalServerError)
		utilities.Log("Dashboard: User not loaded. User Type assert failed.")
		return
	}

	// Type assert failed
	utilities.JsonResponse("Dashboard: Welcome User "+ user.Email, w, http.StatusOK)
	utilities.Log("Dashboard: Welcome User: "+ user.Email)
}
