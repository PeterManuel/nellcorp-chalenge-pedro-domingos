package handlers

import (
	"encoding/json"
	"mybankapi/database"
	"mybankapi/forms"
	"mybankapi/models"
	"net/http"
)

// GetAccounts retrieves all accounts from the database and returns them as JSON.
func Transferir(w http.ResponseWriter, r *http.Request) {
	db, _ := database.ConnectToDatabase()
	var transferencia forms.Transferencia

	// Parse the request body into a Deposito struct.
	err := json.NewDecoder(r.Body).Decode(&transferencia)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	erros, id := models.CreateTransacao(db, "transferencia")

	if erros == nil {
		models.UpdateConta(db, transferencia.IdContaEmissora, -transferencia.Montante)
		e1 := models.InserirLevantamento(db, id, transferencia.IdContaEmissora, transferencia.Montante)
		if e1 != nil {
			http.Error(w, "Falha ao efectuar o Levantamento", http.StatusInternalServerError)
			return
		}
		models.UpdateConta(db, transferencia.IdContaReceptora, transferencia.Montante)
		e2 := models.InserirDeposito(db, id, transferencia.IdContaReceptora, transferencia.Montante)
		if e2 != nil {
			http.Error(w, "Falha ao efectuar o Deposito", http.StatusInternalServerError)
			return
		}

		e3 := models.InserirTransferencia(db, id, transferencia.IdContaEmissora, transferencia.IdContaReceptora, transferencia.Montante)
		if e3 != nil {
			http.Error(w, "Falha ao efectuar a transferencia", http.StatusInternalServerError)
			return
		}

	}

	// Respond with the inserted data, including the ID.
	var msg models.Mensagem
	msg.Descricao = "Transferencia efectuada com sucesso"
	msg.Estado = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
