package forms

import (
	_ "github.com/lib/pq"
)

type Transferencia struct {
	IdContaEmissora  int     `json:"idcontaemissora"`
	IdContaReceptora int     `json:"idcontareceptpra"`
	Montante         float32 `json:"montante"`
}
