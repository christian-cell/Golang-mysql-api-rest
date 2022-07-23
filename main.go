package main

import (
	"log"
	"net/http"

	"mi-api-golang/controllers"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

func main() {

	/* logs cleaner docker logs --since 30s -f <container_name_or_id> */

	r := mux.NewRouter()

	r.HandleFunc("/clientes", controllers.GetClientes).Methods("GET")
	r.HandleFunc("/cliente/{id}", controllers.GetClienteById).Methods("GET")
	r.HandleFunc("/nuevo", controllers.CreateCliente).Methods("POST")
	r.HandleFunc("/actualiza/{id}", controllers.UpdateCliente).Methods("PUT")
	r.HandleFunc("/elimina/{id}", controllers.DeleteCliente).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

	// router := mux.NewRouter().

}
