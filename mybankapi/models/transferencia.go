package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Transferencia struct {
	IdContaEmissora  int     `json:"idcontaemissora"`
	IdContaReceptora int     `json:"idcontareceptpra"`
	Montante         float32 `json:"montante"`
}

func InserirTransferencia(db *sql.DB, idtransacao int, idContaE int, idContaR int, montante float32) error {
	query := `INSERT INTO transferencia (idtransacao,idcontaemissora,idcontareceptora,montante)
              VALUES ($1, $2,$3,$3) RETURNING id`

	_, err := db.Exec(query, idtransacao, idContaE, idContaR, montante)
	return err
}
