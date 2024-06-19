package cyclone

import (
	"math/big"
)

type TxMessageRequest struct {
	Hash        string   `json:"hash"`
	Nonce       *big.Int `json:"nonce"`
	VM          string   `json:"vm"`
	Sender      string   `json:"sender"`
	Message     string   `json:"message"`
	Signature   string   `json:"signature"`
	FeeCurrency string   `json:"feeCurrency"`
}

type vmMessage1 struct {
	Method string `json:"method"`
	Data   string `json:"data"`
}
