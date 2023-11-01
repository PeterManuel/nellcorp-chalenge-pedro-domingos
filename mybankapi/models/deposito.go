package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Deposito struct {
	ID          int     `json:"id"`
	Tipo        string  `json:"tipo"`
	Estado      string  `json:"estado"`
	IdConta     int     `json:"idconta"`
	IdTransacao int     `json:"idtransacao"`
	Montante    float32 `json:"montante"`
}

func InserirDeposito(db *sql.DB, idtransacao int, idConta int, montante float32) error {
	query := `INSERT INTO deposito (idtransacao,idconta,montante)
              VALUES ($1, $2,$3) RETURNING id`

	_, err := db.Exec(query, idtransacao, idConta, montante)
	return err
}

// InsertDeposito inserts data into the deposito table.

func DepositoPorConta(db *sql.DB, idconta int) ([]Deposito, error) {
	query := `
	SELECT d.id, d.idconta, d.idtransacao, d.montante, t.tipo, t.estado
	FROM deposito d
	JOIN transacao t ON d.idtransacao = t.id
	WHERE d.idconta = $1
	`

	rows, err := db.Query(query, idconta)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	depositos := []Deposito{}

	// Iterate through the rows and populate the deposits slice
	for rows.Next() {
		var deposito Deposito

		err := rows.Scan(&deposito.ID, &deposito.IdConta, &deposito.IdTransacao, &deposito.Montante, &deposito.Tipo, &deposito.Estado)
		if err != nil {
			log.Fatal(err)
		}

		depositos = append(depositos, deposito)
	}
	return depositos, nil
}

func DepositoPorTransacao(db *sql.DB, idtransacao int) (int, float32, error) {
	query := `
	SELECT d.idconta, d.montante
	FROM deposito d
	JOIN transacao t ON d.idtransacao = t.id
	WHERE d.idtransacao = $1
	`

	rows, err := db.Query(query, idtransacao)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through the rows and populate the deposits slice
	for rows.Next() {
		var deposito Deposito

		err := rows.Scan(&deposito.IdConta, &deposito.Montante)
		if err != nil {
			log.Fatal(err)
		}
		return deposito.IdConta, deposito.Montante, nil
		//depositos = append(depositos, deposito)
	}
	return -1, -1, nil
}
