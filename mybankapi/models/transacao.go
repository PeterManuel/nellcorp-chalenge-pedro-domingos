package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Transacao struct {
	ID   int    `json:"id"`
	Tipo string `json:"tipo"`

	Estado string `json:"estado"`
}

type TransactionList struct {
	Transferencias []Transferencia
	Levantamentos  []Levantamento
	Depositos      []Deposito
}

func CreateTransacao(db *sql.DB, tipo string) (error, int) {
	query := `INSERT INTO transacao (tipo)
              VALUES ($1) RETURNING id`
	var id int

	err := db.QueryRow(query, tipo).Scan(&id)
	return err, id
}

func GettransacaoByID(db *sql.DB, id int) (Transacao, error) {
	query := "SELECT * FROM transacao WHERE id = $1"

	var transacao Transacao
	err := db.QueryRow(query, id).Scan(&transacao.ID, &transacao.Tipo)
	if err != nil {
		return transacao, err
	}
	return transacao, nil
}

func Updatetransacao(db *sql.DB, id int) error {
	query := `UPDATE transacao
              SET estado='cancelado'
              WHERE id=$1`

	_, err := db.Exec(query, id)
	return err
}

func Deletetransacao(db *sql.DB, id int) error {
	query := "DELETE FROM transacao WHERE id = $1"

	_, err := db.Exec(query, id)
	return err
}
