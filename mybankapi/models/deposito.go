package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Deposito struct {
	ID          int     `json:"id"`
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

/*
func GettransacaoByID(db *sql.DB, id int) (Transacao, error) {
	query := "SELECT * FROM transacao WHERE id = $1"

	var transacao Transacao
	err := db.QueryRow(query, id).Scan(&transacao.ID, &transacao.Tipo, &transacao.Tempo)
	if err != nil {
		return transacao, err
	}
	return transacao, nil
}

func Updatetransacao(db *sql.DB, transacao Transacao) error {
	query := `UPDATE transacao
              SET type=$2, amount=$3, sender_account=$4, receiver_account1=$5, receiver_account2=$6, tempo=$7
              WHERE id=$1`

	_, err := db.Exec(query, transacao.ID, transacao.Tipo, transacao.Montante)
	return err
}

func Deletetransacao(db *sql.DB, id int) error {
	query := "DELETE FROM transacao WHERE id = $1"

	_, err := db.Exec(query, id)
	return err
}
*/
