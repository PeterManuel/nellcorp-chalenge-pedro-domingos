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

func Reembolsar(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idContaStr := vars["idTransacao"]

	// Convert the idConta string to a float64
	idConta, err := strconv.Atoi(idContaStr)
	if err != nil {
		http.Error(w, "Invalid idConta format", http.StatusBadRequest)
		return
	}
	db, _ := database.ConnectToDatabase()

	idconta, montante, _ := models.DepositoPorTransacao(db, idConta)

	models.UpdateConta(db, idconta, -montante)

	models.Updatetransacao(db, idConta)

	var msg models.Mensagem
	msg.Descricao = "Reembolso efectuada com sucesso"
	msg.Estado = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
