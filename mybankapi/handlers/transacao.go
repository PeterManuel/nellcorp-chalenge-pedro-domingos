package handlers

import (
	"encoding/json"
	"mybankapi/database"

	"mybankapi/models"
	"net/http"
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
