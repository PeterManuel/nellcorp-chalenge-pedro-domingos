package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Consulta struct {
	Saldo float64 `json:"saldo"`
}

func ConsultarSaldo(db *sql.DB, id int) (Consulta, error) {
	query := "SELECT saldo FROM conta WHERE id = $1"

	var consulta Consulta
	err := db.QueryRow(query, id).Scan(&consulta.Saldo)
	if err != nil {
		return Consulta{}, err
	}
	return consulta, nil
}
