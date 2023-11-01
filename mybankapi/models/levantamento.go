package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Levantamento struct {
	ID          int     `json:"id"`
	Tipo        string  `json:"tipo"`
	Estado      string  `json:"estado"`
	IdConta     int     `json:"idconta"`
	IdTransacao int     `json:"idtransacao"`
	Montante    float32 `json:"montante"`
}

func InserirLevantamento(db *sql.DB, idtransacao int, idConta int, montante float32) error {
	query := `INSERT INTO levantamento (idtransacao,idconta,montante)
              VALUES ($1, $2,$3) RETURNING id`

	_, err := db.Exec(query, idtransacao, idConta, montante)
	return err
}

func LevantamentoPorConta(db *sql.DB, idconta int) ([]Levantamento, error) {
	query := `
	SELECT d.id, d.idconta, d.idtransacao, d.montante, t.tipo, t.estado
	FROM levantamento d
	JOIN transacao t ON d.idtransacao = t.id
	WHERE d.idconta = $1
	`

	rows, err := db.Query(query, idconta)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	levantamentos := []Levantamento{}

	// Iterate through the rows and populate the deposits slice
	for rows.Next() {
		var levantamento Levantamento

		err := rows.Scan(&levantamento.ID, &levantamento.IdConta, &levantamento.IdTransacao, &levantamento.Montante, &levantamento.Tipo, &levantamento.Estado)
		if err != nil {
			log.Fatal(err)
		}

		levantamentos = append(levantamentos, levantamento)
	}
	return levantamentos, nil
}
