package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

var db *sql.DB

var server = "192.168.0.198"
var port = 1433
var user = "sa"
var password = "jumiadmin"
var database = "dev"

func main() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;encrypt=disable",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	// Read employees
	count, err := ReadEmployees()
	if err != nil {
		log.Fatal("Error reading Employees: ", err.Error())
	}
	fmt.Printf("Read %d row(s) successfully.\n", count)

}

// ReadEmployees reads all employee records
func ReadEmployees() (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	//var nombre string
	//var liquidacion int
	//var detraer float64


	//fmt.Print("Ingrese numero de liquidacion:")
	//fmt.Scan(&liquidacion)
	//fmt.Print("Ingrese importe a detraer:")
	//fmt.Scan(&detraer)

	//fmt.Print("Ingrese nombre:")
	//fmt.Scan(&nombre)


	tsql := fmt.Sprintf("exec prueba3")
	//exec txtAfip_v2 @Periodo = '30/06/2019', @NumeroLiq = 1462, @ImporteDetraer = 4000.32



	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var name string
		var id string

		// Get values from row.
		err := rows.Scan(&id, &name)
		if err != nil {
			return -1, err
		}

		fmt.Printf("ID: %s, Name: %s", id, name)
		count++
	}

	return count, nil
}
