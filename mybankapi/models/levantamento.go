package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Levantamento struct {
	ID          int `json:"id"`
	IdConta     int `json:"idconta"`
	IdTransacao int `json:"idtransacao"`
}

func InserirLevantamento(db *sql.DB, idtransacao int, idConta int, montante float32) error {
	query := `INSERT INTO levantamento (idtransacao,idconta,montante)
              VALUES ($1, $2,$3) RETURNING id`

	_, err := db.Exec(query, idtransacao, idConta, montante)
	return err
}
