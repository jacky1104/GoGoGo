package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

var databaseUrl = "172.16.95.101:3306"
var snDatabaseUrl = "172.16.95.179:3306"
var userName = "forticrm"
var password = "forticrm"
var databaseName = "portal"

func TestMysql(t *testing.T) {

	db, err := sql.Open("mysql", userName+":"+password+"@tcp("+databaseUrl+")/"+databaseName)
	defer db.Close()
	if err != nil {
		fmt.Print("connect to database error:", databaseUrl)
		return
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	rows, err := db.Query("select sn from portal.device limit 500")
	if err != nil {
		fmt.Print("query error:", err)
		return
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var snSlice []string
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for _, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
				snSlice = append(snSlice, value)
			}
		}

		fmt.Println("-----------------------------------")
	}

	labelMap := make(map[string]string)
	for _, value := range snSlice {
		snDb, err := sql.Open("mysql", userName+":"+password+"@tcp("+snDatabaseUrl+")/"+value)
		defer snDb.Close()
		if err != nil {
			fmt.Println("sn:", value, ",error")
			continue
		}
		result, err := snDb.Query("select label from stringMapping where attribute = 'subtype' and (type & 512 = 512) limit 500")
		if err != nil {
			fmt.Println("sn:", value, ",query error,", err)
			continue
		}
		for result.Next() {
			var labelTemp string
			result.Scan(&labelTemp)
			fmt.Println("labelTemp:", labelTemp)
			labelMap[labelTemp] = labelTemp
		}
	}

	for _, value := range labelMap {

		fmt.Println(value)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

var dbUrl = "192.168.223.179:3306"
var dbUserName = "forticrm"
var dbUserPassword = "flzx3kc"

func TestQuerySubType(t *testing.T) {
	db, err := sql.Open("mysql", dbUserName+":"+dbUserPassword+"@tcp("+dbUrl+")/?allowOldPasswords=true")
	defer db.Close()
	if err != nil {
		fmt.Print("connect to database error:", databaseUrl)
		return
	}

	//err = db.Ping()
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}

	rows, err := db.Query("select table_schema from information_schema.tables where table_schema like 'F%' limit 5000")
	if err != nil {
		fmt.Print("query error:", err)
		return
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var snMap = make(map[string]string)
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for _, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
				snMap[value]=value
			}
		}

		fmt.Println("-----------------------------------")
	}

	labelMap := make(map[string]string)
	for _, value := range snMap {
		snDb, err := sql.Open("mysql", dbUserName+":"+dbUserPassword+"@tcp("+dbUrl+")/"+value + "?allowOldPasswords=true")
		if err != nil {
			fmt.Println("sn:", value, ",error")
			continue
		}
		result, err := snDb.Query("select  label from stringMapping where attribute = 'subtype' and (type & 512 ) = 512")
		if err != nil {
			fmt.Println("sn:", value, ",query error,", err)
			continue
		}
		for result.Next() {
			var labelTemp string
			result.Scan(&labelTemp)
			if labelTemp == "virus"{
				fmt.Println("sn,labelTemp:", value,labelTemp)
			}
			labelMap[labelTemp] = labelTemp
		}
		snDb.Close()
	}

	for _, value := range labelMap {

		fmt.Println(value)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
