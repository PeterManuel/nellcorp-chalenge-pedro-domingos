package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Conta struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Saldo float64 `json:"saldo"`
	// Additional account attributes, e.g., account type, etc.
}

func CriarConta(db *sql.DB, conta Conta) error {
	query := `INSERT INTO conta (id,nome,saldo)
              VALUES ($1, $2, $3) RETURNING id`

	_, err := db.Exec(query, conta.Nome, conta.Saldo)
	return err
}

func GetAccountByID(db *sql.DB, id int) (Conta, error) {
	query := "SELECT id, owner_id, owner_name, balance FROM account WHERE id = $1"

	var conta Conta
	err := db.QueryRow(query, id).Scan(&conta.ID, &conta.Nome, &conta.Saldo)
	if err != nil {
		return Conta{}, err
	}
	return conta, nil
}

func UpdateConta(db *sql.DB, idconta int, montante float32) (sql.Result, error) {
	query := `UPDATE conta
              SET  saldo=saldo+$1
              WHERE id=$2`

	result, err := db.Exec(query, montante, idconta)
	return result, err
}

func DeleteAccount(db *sql.DB, id int) error {
	query := "DELETE FROM conta WHERE id = $1"

	_, err := db.Exec(query, id)
	return err
}
