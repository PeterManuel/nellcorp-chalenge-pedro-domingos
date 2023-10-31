package forms

import (
	_ "github.com/lib/pq"
)

type Levantamento struct {
	IdConta     int     `json:"idconta"`
	IdTransacao float32 `json:"idtransacao"`
	Montante    float32 `json:"montante"`
}
