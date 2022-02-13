package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/ingjeffer/crud-system-develoteca/constants"
	"github.com/ingjeffer/crud-system-develoteca/routers"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	router := mux.NewRouter()

	// http.HandleFunc("/", Inicio)
	router.HandleFunc("/", Inicio2)
	router.HandleFunc("/employee", routers.InsertEmployee).Methods(constants.POST)

	// handler := cors.AllowAll().Handler(router)

	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8087", router)
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hola a todos JMC")
	// conexionEstablecida := ConnectionDB()
	// query := "INSERT INTO empleados(nombre, correo) VALUES ('jefferson', 'correo1@correo.com')"
	// insertarRegistros, err := conexionEstablecida.Prepare(query)

	// if err != nil {
	// 	panic(err.Error())
	// }
	// insertarRegistros.Exec()
	// log.Println("hubo conexión")

	plantillas.ExecuteTemplate(w, "inicio", nil)
}

// func Inicio(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w, "Hola a todos JMC")
// 	conexionEstablecida := ConnectionDB()
// 	query := "INSERT INTO empleados(nombre, correo) VALUES ('jefferson', 'correo1@correo.com')"
// 	insertarRegistros, err := conexionEstablecida.Prepare(query)

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	insertarRegistros.Exec()
// 	log.Println("hubo conexión")

// 	plantillas.ExecuteTemplate(w, "inicio", nil)
// }

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

}

func Inicio2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== Inicio2 ===")
}
