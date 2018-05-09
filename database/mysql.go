package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/writetosalman/go-rest-api-boilerplate/utilities"
	"github.com/writetosalman/go-rest-api-boilerplate/config"
)

var dbConn *sql.DB = nil

/**
 * Check the connection and make it if it is lost or not opened yet.
 * @return	error
 */
func checkConnection() error {
	// If already connected then OK
	if dbConn != nil {
		return nil
	}

	// Else try to connect
	conn, err := openDatabase()

	// If error then return with it
	if err != nil {
		return err
	}

	// Else set the connection
	dbConn = conn

	// Tell that it is OK
	return nil
}

/**
 * openDatabase function opens a MySql database connection
 * @param 	*sql.DB
 * @param	Error
 */
func openDatabase() (*sql.DB, error) {

	db, err 	:= sql.Open("mysql", config.Getenv("CONNECTION_STRING")) 	// user:password@tcp(host)/dbname
	if err != nil {
		utilities.Log("MySQL: DB connection failed - " + err.Error())
		return nil, err
	}
	utilities.Log("MySQL: DB connection made")
	return db, nil
}

/**
 * SqlInsert function runs an insert query to MySql
 * @param 	string
 * @param	[]string
 */
func SqlInsert(sqlQuery string, sqlArgument []string) (error, int) {

	err := checkConnection()
	if err != nil {
		utilities.Log("MySQL: DB Insert failed - " + err.Error())
		return err, 0
	}
	//	defer db.Close()

	// Insert
	stmt, err := dbConn.Prepare(sqlQuery)					// "INSERT userinfo SET username=?,departname=?,created=?"
	if err != nil {
		utilities.Log("MySQL: DB Insert failed. Prepare failed: " + err.Error())
		return err, 0
	}

	// Execute Statement
	sqlInterface := utilities.StringArrayToInterface(sqlArgument)

	// Unpacking array | https://stackoverflow.com/questions/17555857/go-unpacking-array-as-arguments
	res, err := stmt.Exec(sqlInterface...)
	if err != nil {
		utilities.Log("MySQL: DB Insert failed. Execute failed: " + err.Error())
		return err, 0
	}

	// Get last insert ID
	id, err := res.LastInsertId()
	if err != nil {
		utilities.Log("MySQL: Last Insert ID failed. " + err.Error())
		return err, 0
	}

	return nil, int(id)
}

/**
 * SqlQuery function runs a select query to MySql
 * @param 	string
 * @param	[]string
 */
func SqlQuery(sqlQuery string, sqlArgument []string) (*sql.Rows, error) {

	err := checkConnection()
	if err != nil {
		utilities.Log("MySQL: DB Query failed. " + err.Error())
		return nil, err
	}
	//	defer db.Close()

	sqlInterface 	:= utilities.StringArrayToInterface(sqlArgument)

	// query
	rows, err 	:= dbConn.Query(sqlQuery, sqlInterface...)
	if err != nil {
		utilities.Log("MySQL: SQL query failed. "+ err.Error())
		return nil, err
	}

	return rows, nil
}

/**
 * SqlUpdate function runs a update query to MySql
 * @param 	string
 * @param	[]string
 */
func SqlUpdate(sqlQuery string, sqlArgument []string) (int, error) {

	err := checkConnection()
	if err != nil {
		utilities.Log("MySQL: DB Update failed. " + err.Error())
		return 0, err
	}
	//	defer db.Close()

	sqlInterface := utilities.StringArrayToInterface(sqlArgument)

	// Prepare
	stmt, err := dbConn.Prepare(sqlQuery)
	if err != nil {
		utilities.Log("MySQL: SQL update prepare failed. "+ err.Error())
		return 0, err
	}

	// Execute
	res, err := stmt.Exec(sqlInterface...)
	if err != nil {
		utilities.Log("MySQL: SQL update execute failed. "+ err.Error())
		return 0, err
	}

	// Get rows affected
	affect, err := res.RowsAffected()
	if err != nil {
		utilities.Log("MySQL: SQL update rows affected failed. "+ err.Error())
		return 0, err
	}

	return int(affect), nil
}


/**
 * SqlDelete function runs a delete query to MySql
 * @param 	string
 * @param	[]string
 */
func SqlDelete(sqlQuery string, sqlArgument []string) (int, error) {

	err := checkConnection()
	if err != nil {
		utilities.Log("MySQL: DB Delete failed. " + err.Error())
		return 0, err
	}
	//	defer db.Close()

	sqlInterface := utilities.StringArrayToInterface(sqlArgument)

	// Prepare
	stmt, err := dbConn.Prepare(sqlQuery)
	if err != nil {
		utilities.Log("MySQL: SQL delete prepare failed. "+ err.Error())
		return 0, err
	}

	// Execute
	res, err := stmt.Exec(sqlInterface...)
	if err != nil {
		utilities.Log("MySQL: SQL delete execute failed. "+ err.Error())
		return 0, err
	}

	// Get rows affected
	affect, err := res.RowsAffected()
	if err != nil {
		utilities.Log("MySQL: SQL delete rows affected failed. "+ err.Error())
		return 0, err
	}

	return int(affect), nil
}

func SqlGetRecord(sqlQuery string, sqlArgument []string, scanned ...interface{}) error {

	// Get the rows
	rows, err := SqlQuery(sqlQuery, sqlArgument)
	if err != nil {
		return err
	}

	// Close Query connection in end
	defer rows.Close()

	// Try to scan the row
	if rows.Next() {
		err = rows.Scan(scanned...)
		if err != nil {
			return err
		}
	}

	// If record not found
	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}
