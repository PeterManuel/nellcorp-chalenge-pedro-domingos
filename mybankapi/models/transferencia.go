package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Transferencia struct {
	ID               int     `json:"id"`
	Tipo             string  `json:"tipo"`
	Estado           string  `json:"estado"`
	IdContaEmissora  int     `json:"idcontaemissora"`
	IdContaReceptora int     `json:"idcontareceptora"`
	Montante         float32 `json:"montante"`
}

func InserirTransferencia(db *sql.DB, idtransacao int, idContaE int, idContaR int, montante float32) error {
	query := `INSERT INTO transferencia (idtransacao,idcontaemissora,idcontareceptora,montante)
              VALUES ($1, $2,$3,$4) RETURNING id`

	_, err := db.Exec(query, idtransacao, idContaE, idContaR, montante)
	return err
}

func TransferenciaPorConta(db *sql.DB, idconta int) ([]Transferencia, error) {
	query := `
	SELECT d.id, d.idcontaemissora,d.idcontareceptora, d.montante, t.tipo, t.estado
	FROM transferencia d
	JOIN transacao t ON d.idtransacao = t.id
	WHERE d.idcontaemissora = $1 or d.idcontareceptora = $1
	`

	rows, err := db.Query(query, idconta)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	transferencias := []Transferencia{}

	// Iterate through the rows and populate the deposits slice
	for rows.Next() {
		var transferencia Transferencia

		err := rows.Scan(&transferencia.ID, &transferencia.IdContaEmissora, &transferencia.IdContaReceptora, &transferencia.Montante, &transferencia.Tipo, &transferencia.Estado)
		if err != nil {
			log.Fatal(err)
		}

		transferencias = append(transferencias, transferencia)
	}
	return transferencias, nil
}
