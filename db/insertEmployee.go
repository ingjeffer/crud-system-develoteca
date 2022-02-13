package db

import (
	"fmt"
	"log"

	"github.com/ingjeffer/crud-system-develoteca/models"
)

func InsertEmployee(employee models.Employee) (bool, error, int64) {
	query := "INSERT INTO employee(name, email) VALUES (?,?)"
	connection := ConnectionDB()
	insert, err := connection.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return false, err, 0
	}
	result, err2 := insert.Exec(employee.Name, employee.Email)
	if err2 != nil {
		log.Println(err.Error())
		return false, err2, 0
	}
	last, _ := result.LastInsertId()
	fmt.Printf("last: %d", last)
	return true, nil, last
}
