package forms

import (
	_ "github.com/lib/pq"
)

type Deposito struct {
	IdConta  int     `json:"idconta"`
	Montante float32 `json:"montante"`
}
