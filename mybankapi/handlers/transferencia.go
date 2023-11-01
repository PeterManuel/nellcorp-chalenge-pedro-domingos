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

	saldo, _ := models.ConsultarSaldo(db, transferencia.IdContaEmissora)

	if saldo.Saldo < float64(transferencia.Montante) {
		var msg models.Mensagem
		msg.Descricao = "Não tem saldo Suficiente"
		msg.Estado = "Error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
		return

	}

	erros, id := models.CreateTransacao(db, "transferencia")

	if erros == nil {
		models.UpdateConta(db, transferencia.IdContaEmissora, -transferencia.Montante)
		e1 := models.InserirLevantamento(db, id, transferencia.IdContaEmissora, transferencia.Montante)
		if e1 != nil {
			var msg models.Mensagem
			msg.Descricao = "Erro Ao Transferir"
			msg.Estado = "Error"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(msg)
			return
		}
		models.UpdateConta(db, transferencia.IdContaReceptora, transferencia.Montante)
		e2 := models.InserirDeposito(db, id, transferencia.IdContaReceptora, transferencia.Montante)
		if e2 != nil {
			var msg models.Mensagem
			msg.Descricao = "Erro Ao Transferir"
			msg.Estado = "Error"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(msg)
			return
		}

		e3 := models.InserirTransferencia(db, id, transferencia.IdContaEmissora, transferencia.IdContaReceptora, transferencia.Montante)
		if e3 != nil {
			var msg models.Mensagem
			msg.Descricao = "Erro Ao Transferir"
			msg.Estado = "Error"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(msg)
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
