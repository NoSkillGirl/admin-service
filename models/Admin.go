package models

import (
	"context"
	"database/sql"
	"fmt"
)

var (
	ctx context.Context
	db  *sql.DB
)

const mySQLHost = "34.93.137.151"

var mySQLConnection = fmt.Sprintf("root:password@tcp(%s)/tour_travel", mySQLHost)

//AddCompany function
func AddCompany(name, ownerName, phoneNo string) (errOccured, duplicateExist bool) {
	ctx := context.Background()
	// mySQLConnection := "root:password@tcp(%s)/tour_travel"
	// mySQLConnection = fmt.Sprintf(mySQLConnection, mySQLHost)
	db, err := sql.Open("mysql", mySQLConnection)
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println(err)
		//panic(err.Error())
		return true, false
	}
	// defer the close till after the main function has finished executing
	defer db.Close()

	var count int
	error := db.QueryRowContext(ctx, "select count(*) from companies where name =?", name).Scan(&count)
	if error != nil {
		return true, false
	}
	if count == 0 {
		addCompanyQuery := `INSERT INTO companies (name, owner_name, phone_no) VALUES ('%s', '%s', '%s')`
		addCompanyQueryString := fmt.Sprintf(addCompanyQuery, name, ownerName, phoneNo)
		fmt.Println(addCompanyQueryString)

		// perform a db.Query insert
		insert, err := db.Query(addCompanyQueryString)

		// if there is an error inserting, handle it
		if err != nil {
			fmt.Println(err)
			return true, false
		}
		// be careful deferring Queries if you are using transactions
		defer insert.Close()
		return false, false
	}
	//error = false, duplicate = true
	return false, true
}

//AddBus function
func AddBus(number string, ac bool, sleeper bool, totalSeats int, source string, destination string, companyID int, departureTime string, arrivalTime string) (errOccured, duplicateExist bool) {
	ctx := context.Background()
	// db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/tour_travel")
	db, err := sql.Open("mysql", mySQLConnection)
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println(err)
		//panic(err.Error())
		return true, false
	}
	// defer the close till after the main function has finished executing
	defer db.Close()
	var count int
	error := db.QueryRowContext(ctx, "select count(*) from bus_details where number=?", number).Scan(&count)
	if error != nil {
		// fmt.Println(error)
		return true, false
	}
	if count == 0 {
		addBusQuery := `INSERT INTO bus_details 
		(number, ac, sleeper, total_seats, source, destination, company_id, departure_time, arrival_time) 
		VALUES ('%s', %t, %t, %d, '%s', '%s', %d, '%s', '%s')`

		addBusQueryString := fmt.Sprintf(addBusQuery, number, ac, sleeper, totalSeats, source, destination, companyID, departureTime, arrivalTime)
		fmt.Println(addBusQueryString)

		insert, err := db.Query(addBusQueryString)
		// if there is an error inserting, handle it
		if err != nil {
			fmt.Println(err)
			return true, false
		}
		// be careful deferring Queries if you are using transactions
		defer insert.Close()
		return false, false
	}
	//error = false, duplicate = true
	return false, true
}
