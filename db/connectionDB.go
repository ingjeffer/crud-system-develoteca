package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "admin"
	Contrasenia := "admin"
	Nombre := "sistema"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}
