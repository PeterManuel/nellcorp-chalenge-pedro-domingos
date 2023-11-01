package handlers

import (
	"encoding/json"
	"mybankapi/database"
	"mybankapi/models"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Consultar(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idContaStr := vars["idConta"]

	// Convert the idConta string to a float64
	idConta, err := strconv.Atoi(idContaStr)
	if err != nil {
		http.Error(w, "Invalid idConta format", http.StatusBadRequest)
		return
	}
	db, _ := database.ConnectToDatabase()

	var conta models.Consulta
	conta, _ = models.ConsultarSaldo(db, idConta)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(conta)
}
