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
var database = "SJJUMI"

func main() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;encrypt=disable",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error al crear la conexión con base de datos: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Conexión establecida!\n")

	// Read employees
	count, err := GenerarTxt()
	if err != nil {
		log.Fatal("Error al obtener los resultados del txt: ", err.Error())
	}
	fmt.Printf("%d fila(s) leída(s) satisfactoriamente.\n", count)

}

// ReadEmployees reads all employee records
func GenerarTxt() (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	var periodo string
	var liquidacion int
	var detraer float64

	fmt.Print("Ingrese periodo(dd/mm/aaaa):")
	fmt.Scan(&periodo)
	fmt.Print("Ingrese liquidacion:")
	fmt.Scan(&liquidacion)
	fmt.Print("Ingrese importe a detraer:")
	fmt.Scan(&detraer)


	//fmt.Println(nombre)


	tsql := fmt.Sprintf("exec txtAfip_v2 '%[1]s' ,%[2]d ,%[3]g\n", periodo, liquidacion, detraer)
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
		var name sql.NullString

		// Get values from row.
		err := rows.Scan(&name)
		if err != nil {
			return -1, err
		}

		//fmt.Printf("ID: %s",name)
		count++
	}

	fmt.Println("Se ha generado correctamente el archivo txt")

	return count, nil
}
