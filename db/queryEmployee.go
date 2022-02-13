package db

import (
	"log"

	"github.com/ingjeffer/crud-system-develoteca/models"
)

func QueryEmployee() (bool, error, []models.Employee) {
	query := "SELECT * FROM employee"
	connection := ConnectionDB()
	registers, err := connection.Query(query)
	if err != nil {
		log.Println(err.Error())
		return false, err, nil
	}
	employees := []models.Employee{}
	for registers.Next() {
		var id int64
		var name, email string
		err = registers.Scan(&id, &name, &email)
		employees = append(employees, models.Employee{
			Email: email,
			Id:    id,
			Name:  name,
		})
	}
	registers.Close()
	return true, nil, employees
}
