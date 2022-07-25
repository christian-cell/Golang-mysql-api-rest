package controllers

import (
	"encoding/json"
	"fmt"
	"mi-api-golang/commons"
	"mi-api-golang/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Saludo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola bienvenido a mi api de golang conectada a mysql")
}

func GetClientes(w http.ResponseWriter, r *http.Request) {
	db, err := commons.GetConnection()

	commons.Catch(err)

	results, err := db.Query("SELECT  nombre , apellido_1 , apellido_2 , edad FROM clientes WHERE true")
	commons.Catch(err)

	var clientes = []models.Clientes{}
	var cliente = models.Clientes{}

	for results.Next() {
		if err := results.Scan(&cliente.Nombre, &cliente.Apellido_1, &cliente.Apellido_2, &cliente.Edad); err != nil {
			commons.Catch(err)
		}

		clientes = append(clientes, cliente)
	}

	commons.JsonRespond(w, http.StatusOK, clientes)
}

func GetClienteById(w http.ResponseWriter, r *http.Request) {

	id_str := mux.Vars(r)

	id, err := strconv.Atoi(id_str["id"]) /* strconv.ParseInt(id_str["id"], 10, 64) */

	commons.Catch(err)

	var cliente models.Clientes

	db, err := commons.GetConnection()

	commons.Catch(err)

	row := db.QueryRow("SELECT nombre , apellido_1 , apellido_2 , edad FROM clientes WHERE id = ?", id)
	err = row.Scan(&cliente.Nombre, &cliente.Apellido_1, &cliente.Apellido_2, &cliente.Edad)
	commons.Catch(err)

	commons.JsonRespond(w, http.StatusOK, cliente)

}

func CreateCliente(w http.ResponseWriter, r *http.Request) {

	var cliente models.Clientes

	err := json.NewDecoder(r.Body).Decode(&cliente)
	commons.Catch(err)

	// json.NewEncoder(w).Encode(cliente)

	db, err := commons.GetConnection()

	commons.Catch(err)

	nombre := &cliente.Nombre
	apellido_1 := &cliente.Apellido_1
	apellido_2 := &cliente.Apellido_2
	edad := &cliente.Edad

	result, err := db.Prepare("INSERT INTO clientes (nombre , apellido_1 , apellido_2 , edad) VALUES (?,?,?,?)")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(&cliente)
	}

	result.Exec(nombre, apellido_1, apellido_2, edad)

	fmt.Fprintf(w, "Cliente was created")
}

func UpdateCliente(w http.ResponseWriter, r *http.Request) {

	var cliente models.Clientes

	err := json.NewDecoder(r.Body).Decode(&cliente)

	params := mux.Vars(r)

	db, err := commons.GetConnection()

	commons.Catch(err)

	cliente_querie, err := db.Prepare("UPDATE clientes SET nombre = ? , apellido_1 = ?, apellido_2 = ? , edad = ? WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	nombre := &cliente.Nombre
	apellido_1 := &cliente.Apellido_1
	apellido_2 := &cliente.Apellido_2
	edad := &cliente.Edad

	_, err = cliente_querie.Exec(nombre, apellido_1, apellido_2, edad, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Clientes with ID = %s was updated", params["id"])

}

func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, err := commons.GetConnection()

	stmt, err := db.Prepare("DELETE FROM clientes WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Cliente with ID = %s was deleted", params["id"])
}
