package cyclone

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cast"
	"go.uber.org/zap"

	"github.com/saiset-co/sai-cyclone-interaction/internal/crypto"
	"github.com/saiset-co/sai-cyclone-interaction/logger"
)

type Cyclone struct {
	ctx     context.Context
	node    string
	wallet  string
	private string
	saiBtc  *crypto.SaiBTC
}

func NewCyclone(ctx context.Context, node, wallet, private string, saiBtc *crypto.SaiBTC) Cyclone {
	return Cyclone{
		ctx:     ctx,
		node:    node,
		private: private,
		wallet:  wallet,
		saiBtc:  saiBtc,
	}
}

func (c *Cyclone) GenerateTransactionMessage(message string) (TxMessageRequest, error) {
	var s = rand.NewSource(time.Now().UnixMicro())

	r := rand.New(s)

	tx := TxMessageRequest{
		Hash:      uuid.New().String(),
		Nonce:     big.NewInt(r.Int63n(1000000)),
		Message:   makeVM1ExecuteMessage(message),
		Sender:    cast.ToString(c.wallet),
		Signature: "",
	}

	signature, err := c.saiBtc.SignMessage(c.ctx, payload(tx), c.private)
	if err != nil {
		logger.Logger.Error("GenerateTransactionMessage", zap.Error(err))
		return tx, err
	}

	tx.Signature = signature

	hash, err := makeHash(tx)
	if err != nil {
		logger.Logger.Error("GenerateTransactionMessage", zap.Error(err))
		return tx, err
	}

	tx.Hash = hash

	return tx, nil
}

func (c *Cyclone) SendTx(url string, data interface{}, response interface{}) error {
	requestBody, err := json.Marshal(data)
	if err != nil {
		logger.Logger.Error("sendQuery", zap.Error(err))
		return err
	}

	req, err := http.NewRequest("POST", c.node+url, bytes.NewBuffer(requestBody))
	if err != nil {
		logger.Logger.Error("sendQuery", zap.Error(err))
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Logger.Error("sendQuery", zap.Error(err))
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		logger.Logger.Error("sendQuery", zap.Error(err))
		return err
	}

	return nil
}

func makeVM1ExecuteMessage(s string) string {
	m := &vmMessage1{
		Method: "execute",
		Data:   s,
	}

	mBytes, _ := json.Marshal(m)

	return string(mBytes)
}

func makeHash(m TxMessageRequest) (string, error) {
	key := m.VM + m.Signature + m.Sender + m.Message + m.Nonce.String() + m.FeeCurrency
	hash, err := CreateHash(key)
	if err != nil {
		logger.Logger.Error("makeHash", zap.Error(err))
		return "", err
	}

	return hash, nil
}

func payload(m TxMessageRequest) string {
	payload := m.VM + m.Sender + m.Message + m.Nonce.String() + m.FeeCurrency

	return base64.StdEncoding.EncodeToString([]byte(payload))
}
