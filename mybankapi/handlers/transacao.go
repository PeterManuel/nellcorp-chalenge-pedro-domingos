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

func ListarTransacao(w http.ResponseWriter, r *http.Request) {

	db, err := database.ConnectToDatabase()

	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT * FROM transacao")
	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var transacoes []models.Transacao
	for rows.Next() {
		var transacao models.Transacao
		err := rows.Scan(&transacao.ID, &transacao.Tipo, &transacao.Estado)
		if err != nil {

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		transacoes = append(transacoes, transacao)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transacoes)
}

func ListarTransacaoPorConta(w http.ResponseWriter, r *http.Request) {

	db, err := database.ConnectToDatabase()

	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	idContaStr := vars["idConta"]

	// Convert the idConta string to a float64
	idConta, err := strconv.Atoi(idContaStr)
	if err != nil {
		var msg models.Mensagem
		msg.Descricao = "Falha ao obter a conta"
		msg.Estado = "Error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
		return
	}

	var transacao models.TransactionList
	transacao.Depositos, _ = models.DepositoPorConta(db, idConta)
	transacao.Levantamentos, _ = models.LevantamentoPorConta(db, idConta)
	transacao.Transferencias, _ = models.TransferenciaPorConta(db, idConta)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transacao)
}
