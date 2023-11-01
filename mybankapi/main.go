package main

import (
	"log"
	"mybankapi/handlers"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	// Create a new router.
	r := mux.NewRouter()

	// Define a route to get all accounts.
	r.HandleFunc("/accounts", handlers.GetAccounts).Methods("GET")

	r.HandleFunc("/depositar", handlers.Depositar).Methods("POST")
	r.HandleFunc("/levantar", handlers.Levantar).Methods("POST")
	r.HandleFunc("/transferir", handlers.Transferir).Methods("POST")
	r.HandleFunc("/reembolsar", handlers.Levantar).Methods("POST")
	r.HandleFunc("/saldo/{idConta}", handlers.Consultar).Methods("GET")
	r.HandleFunc("/transacoes/{idConta}", handlers.ListarTransacaoPorConta).Methods("GET")

	r.HandleFunc("/transacao", handlers.ListarTransacao).Methods("GET")
	r.HandleFunc("/depositos", handlers.ListarDepositos).Methods("GET")

	// Start the server.
	log.Fatal(http.ListenAndServe(":8080", r))
}
