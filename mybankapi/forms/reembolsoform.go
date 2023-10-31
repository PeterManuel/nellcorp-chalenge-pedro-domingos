package forms

import (
	_ "github.com/lib/pq"
)

type Reembolso struct {
	IdTransacao int `json:"idtransacao"`
}
