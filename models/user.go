package models

import (
	"errors"

	"github.com/writetosalman/go-rest-api-boilerplate/utilities"
	"github.com/writetosalman/go-rest-api-boilerplate/database"
)

type User struct {
	UserID 				int
	Email				string
	FirstName 			string
	LastName 			string
	PasswordHash 		string
	IsGoogle2faEnabled	bool
	Google2faSecret		string // TODO add support for 2FA
}


// GetUserRecordByEmail returns map of user record
func GetUserByEmail(email string) (User, error) {

	var user 	User

	if email == "" {
		return user, errors.New("Unable to get user. Email is empty")
	}

	sqlQuery 		:= "SELECT id, first_name, last_name, password FROM users where email=?"
	sqlArgument 		:= []string{email}
	rows, err 		:= database.SqlQuery(sqlQuery, sqlArgument)

	if err != nil {
		utilities.Log("Unable to get user by email from database")
		return user, errors.New("Unable to get user by email from database")
	}

	defer rows.Close() 	// Close Query connection in end
	for rows.Next() {

		// TODO | Salman Get more fields from db like Google 2FA

		err = rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.PasswordHash)
		utilities.Log("Login found in database: " + user.FirstName + " " + user.LastName)
		if err != nil {
			return user, errors.New("Unable to get user record:" + err.Error())
		}

		// Set email passed
		user.Email = email

		return user, nil
	}

	// If record not found
	err = rows.Err()
	if err != nil {
		utilities.Log("User not found")
		return user, errors.New("User not found in db:" + err.Error())
	}
	utilities.Log("User not found")
	return user, errors.New("User not found")
}


// ValidateUserID function checks if user is logged in
func GetUserByID(userID string) (*User, error) {

	if userID == "" {
		utilities.Log("Cannot find user. UserID is zero")
		return nil, errors.New("Cannot find user. UserID is zero")
	}

	sqlQuery 		:= "select email, first_name, last_name, password from users where id=?"
	sqlArgument 		:= []string{ userID }
	rows, err 		:= database.SqlQuery(sqlQuery, sqlArgument)

	if err != nil {
		utilities.Log("Unable to get user by ID from database. ID: "+ userID)
		return nil, errors.New("Unable to get user by ID from database.")
	}

	defer rows.Close() 	// Close Query connection in end
	for rows.Next() {

		var user User

		// TODO | Salman Get more fields from db like Google 2FA

		err = rows.Scan(&user.Email, &user.FirstName, &user.LastName, &user.PasswordHash)
		if err != nil {
			utilities.Log("User NOT found in database")
			return nil, errors.New("Unable to get user record:" + err.Error())
		}
		utilities.Log("User found in database: " + user.FirstName + " " + user.LastName)

		// Set ID passed as argument
		user.UserID = utilities.StringToInt(userID)

		// Return User
		return &user, nil
	}

	// If record not found
	err = rows.Err()
	if err != nil {
		utilities.Log("UserID not found")
		return nil, errors.New("UserID not found")
	}

	return nil, errors.New("UserID validated")
}
