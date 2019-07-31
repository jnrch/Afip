package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
)

var server = "192.168.0.198"
var port = 1433
var user = "sa"
var password = "jumiadmin"
var database = "dev"

// Define a User model struct
type Prueba struct {
	gorm.Model
	Nombre string
	Apellido string
}


// Read and print all the tasks
func ReadAllTasks(db *gorm.DB){
	var pruebas []Prueba
	db.Find(&pruebas)

	for _, prueba := range pruebas{
		db.Model(&prueba)
		fmt.Printf("%s %s's tasks:\n", prueba.Nombre, prueba.Apellido)
	}
}

func main() {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;encrypt=disable",
		server, user, password, port, database)
	db, err := gorm.Open("mssql", connectionString)

	if err != nil {
		log.Fatal("Failed to create connection pool. Error: " + err.Error())
	}
	defer db.Close()

	// Read
	fmt.Println("\nReading all the tasks...")
	ReadAllTasks(db)


}

