package handlers

import (
	"encoding/json"
	"mybankapi/database"
	"mybankapi/forms"
	"mybankapi/models"
	"net/http"
)

// GetAccounts retrieves all accounts from the database and returns them as JSON.
func Levantar(w http.ResponseWriter, r *http.Request) {
	db, _ := database.ConnectToDatabase()
	var levantamento forms.Levantamento

	// Parse the request body into a Deposito struct.
	err := json.NewDecoder(r.Body).Decode(&levantamento)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res, e := models.UpdateConta(db, levantamento.IdConta, -levantamento.Montante)
	n, _ := res.RowsAffected()
	if e != nil || n == 0 {
		http.Error(w, "Falha ao efectuar o Levantamento", http.StatusInternalServerError)
		return
	}

	erros, id := models.CreateTransacao(db, "levantamento")
	if erros == nil {
		e1 := models.InserirLevantamento(db, id, levantamento.IdConta, levantamento.Montante)
		if e1 != nil {
			http.Error(w, "Falha ao efectuar o Levantamento", http.StatusInternalServerError)
			return
		}

	}

	// Respond with the inserted data, including the ID.
	var msg models.Mensagem
	msg.Descricao = "Levantamento efectuado com sucesso"
	msg.Estado = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func ListarLevantamentos(w http.ResponseWriter, r *http.Request) {

	db, err := database.ConnectToDatabase()

	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT * FROM levantamento")
	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var levantamentos []models.Levantamento
	for rows.Next() {
		var levantamento models.Levantamento
		err := rows.Scan(&levantamento.ID, &levantamento.IdConta, &levantamento.IdTransacao)
		if err != nil {

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		levantamentos = append(levantamentos, levantamento)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(levantamentos)
}
