//  Generate RSA signing files via shell (adjust as needed):
//
//  $ openssl genrsa -out server.key 2048
//  $ openssl rsa -in server.key -pubout > server.pub

package services

import (
	"crypto/rsa"
	"io/ioutil"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"errors"
	"fmt"

	"github.com/writetosalman/go-rest-api-boilerplate/utilities"
)

const (
	privateKeyPath 	= "keys/server.key"
	pubKeyPath  	= "keys/server.pub"
)

var (
	verifyKey 	*rsa.PublicKey
	signKey   	*rsa.PrivateKey
)

// Initialise Public & Private Keys
func initialiseKeys() {
	signBytes, err 		:= ioutil.ReadFile(privateKeyPath)
	utilities.CheckError(err)

	signKey, err 		= jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	utilities.CheckError(err)

	verifyBytes, err 	:= ioutil.ReadFile(pubKeyPath)
	utilities.CheckError(err)

	verifyKey, err 		= jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	utilities.CheckError(err)
}

// This function creates HASH of the password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	utilities.Log("Password Hashed: " + string(bytes))
	return string(bytes), err
}

// This function checks HASH of the password with the database hash saved
func CheckPasswordHash(password, hash string) bool {

	// TODO Salman Fix this and make it work
	utilities.Log("Password ["+ password +"] ["+ hash +"]")
	return true

	//err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	//utilities.Log("Password Check: " + err.Error())
	//return true || err == nil // Just for testing
}

// GetJwtClaims create claims and return for use in JWT
func GetJwtClaims(userID int, email string) (jwt.MapClaims) {

	claims 				:= make(jwt.MapClaims)
	claims["userID"] 	= utilities.IntToString(userID)
	claims["email"] 	= email
	claims["exp"] 		= time.Now().Add(time.Hour * 4).Unix()

	return claims
}

// GetJwtToken creates JWT token from UserID & email address
func GetJwtToken(userID int, email string) (string, error) {

	// Initialise Keys
	initialiseKeys()

	// Create token
	token 			:= jwt.New(jwt.SigningMethodRS256)

	// Set token claims
	token.Claims 		= GetJwtClaims(userID, email)

	// Sign token with key
	tokenString, err 	:= token.SignedString(signKey)
	if err != nil {
		utilities.Log("Failed to sign token")
		return "", errors.New("Failed to sign token")
	}

	return tokenString, nil
}

// ValidateJwtToken as name suggests validates JWT Token
func VerifyJwtToken(token string) (map[string]string, bool) {

	// Initialise Keys
	initialiseKeys()

	// If the token is empty
	if token == "" {
		return map[string]string{"error":"Token is empty"}, false
	}

	// Now parse the token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			msg := fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, msg
		}
		// Don't know what it is doing so commented it | return a.encryptionKey, nil
		return verifyKey, nil
	})

	// Parsing failed
	if err != nil {
		utilities.Log("Error parsing token:", err.Error())
		return map[string]string{"error":"Error parsing token: "+ err.Error()}, false
	}

	// Check token is valid
	if parsedToken != nil && parsedToken.Valid {
		arrTokenClaims 	:= parsedToken.Claims.(jwt.MapClaims)

		strEmail 	:= utilities.InterfaceToString(arrTokenClaims["email"])
		strUserID 	:= utilities.InterfaceToString(arrTokenClaims["userID"])

		utilities.Log("Success in parsing token")
		return map[string]string{"email":strEmail, "userID":strUserID}, true
	}
	return map[string]string{"error":"Error: Token rejected"}, false
}

// ValidateLoginCredentials checks if username and password is correct
// If password is verified then an integer other than zero is returned
func VerifyUserPassword(password string, passwordHash string) (string, bool) {

	if CheckPasswordHash(password, passwordHash) {
		utilities.Log("Password verified.")
		return "Password verified", true
	}

	utilities.Log("Incorrect Password")
	return "Incorrect Password", false
}
