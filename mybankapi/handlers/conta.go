package handlers

import (
	"encoding/json"
	"log"
	"mybankapi/database"
	"mybankapi/models"
	"net/http"
)

// GetAccounts retrieves all accounts from the database and returns them as JSON.
func GetAccounts(w http.ResponseWriter, r *http.Request) {

	db, err := database.ConnectToDatabase()

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT * FROM conta")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var contas []models.Conta
	for rows.Next() {
		var conta models.Conta
		err := rows.Scan(&conta.ID, &conta.Nome, &conta.Saldo)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		contas = append(contas, conta)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contas)
}
