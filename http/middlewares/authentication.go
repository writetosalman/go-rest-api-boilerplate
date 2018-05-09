package middlewares

import (
	"net/http"
	"strings"
	"fmt"

	"github.com/gorilla/context"
	AuthenticateService "github.com/writetosalman/go-rest-api-boilerplate/services"
	"github.com/writetosalman/go-rest-api-boilerplate/utilities"
	"github.com/writetosalman/go-rest-api-boilerplate/models"
)


func Authentication(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get token from the Authorization header | format: Authorization: Bearer
		var token string
		tokens, isOk := r.Header["Authorization"]
		if isOk && len(tokens) >= 1 {
			token = tokens[0]
			token = strings.TrimPrefix(token, "Bearer ")
		}

		arrParsedToken, isError := AuthenticateService.VerifyJwtToken(token)
		if isError == false {
			utilities.Log("Authentication: "+ arrParsedToken["error"])
			utilities.JsonResponse(arrParsedToken["error"], w, http.StatusUnauthorized)
			return;
		}

		// Token validated
		utilities.Log("Token validated for ("+arrParsedToken["userID"]+") email: "+ arrParsedToken["email"] +"")

		// Check if userID is not empty
		if arrParsedToken["userID"] == "" {
			utilities.Log("Authentication: Unauthorized - unable to validate token")
			utilities.JsonResponse("Unauthorized: Unable to validate token", w, http.StatusUnauthorized)
			return;
		}

		// Get User record from db
		User, err	:= models.GetUserByID( arrParsedToken["userID"] )
		if err != nil {
			utilities.Log("Authentication: User not found")
			utilities.JsonResponse("Unauthorized: User not found", w, http.StatusUnauthorized)
			return
		}

		// User found in database
		utilities.Log("User found in database. Now set context and return from this middleware")


		// Pass user information to context so that it is available in controllers
		// More info: https://golang.org/pkg/context/
		context.Set(r, "User",  *User)
		fmt.Println(*User)

		// Success
		next.ServeHTTP(w, r)
	})
}
