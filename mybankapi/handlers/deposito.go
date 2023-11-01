package handlers

import (
	"encoding/json"
	"mybankapi/database"
	"mybankapi/forms"
	"mybankapi/models"
	"net/http"
)

// GetAccounts retrieves all accounts from the database and returns them as JSON.
func Depositar(w http.ResponseWriter, r *http.Request) {
	db, _ := database.ConnectToDatabase()
	var deposito forms.Deposito

	// Parse the request body into a Deposito struct.
	err := json.NewDecoder(r.Body).Decode(&deposito)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res, e := models.UpdateConta(db, deposito.IdConta, deposito.Montante)
	n, _ := res.RowsAffected()
	if e != nil || n == 0 {
		var msg models.Mensagem
		msg.Descricao = "Falha ao efectuar o Deposito"
		msg.Estado = "Error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
		return
	}

	erros, id := models.CreateTransacao(db, "deposito")
	if erros == nil {

		e1 := models.InserirDeposito(db, id, deposito.IdConta, deposito.Montante)
		if e1 != nil {
			var msg models.Mensagem
			msg.Descricao = "Falha ao efectuar o Deposito"
			msg.Estado = "Error"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(msg)
			return
		}

	}

	// Respond with the inserted data, including the ID.
	var msg models.Mensagem
	msg.Descricao = "Desposito efectuado com sucesso"
	msg.Estado = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func ListarDepositos(w http.ResponseWriter, r *http.Request) {

	db, err := database.ConnectToDatabase()

	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT * FROM deposito")
	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var depositos []models.Deposito
	for rows.Next() {
		var deposito models.Deposito
		err := rows.Scan(&deposito.ID, &deposito.IdConta, &deposito.IdTransacao, &deposito.Montante)
		if err != nil {

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		depositos = append(depositos, deposito)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(depositos)
}
