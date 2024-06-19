package internal

import (
	"encoding/json"

	validator "github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/saiset-co/sai-cyclone-interaction/logger"
	saiService "github.com/saiset-co/sai-service/service"
)

func (is *InternalService) NewHandler() saiService.Handler {
	return saiService.Handler{
		"send_tx": saiService.HandlerElement{
			Name:        "Send transaction",
			Description: "Send new raw transaction",
			Function: func(data, meta interface{}) (interface{}, int, error) {
				return is.sendRawTransaction(data)
			},
		},
	}
}

func (is *InternalService) sendRawTransaction(data interface{}) (interface{}, int, error) {
	var request = new(CreateTxRequest)
	var response = new(CreateTxResponse)

	dataJson, err := json.Marshal(data)
	if err != nil {
		logger.Logger.Error("sendRawTransaction", zap.Error(err))
		return nil, 500, err
	}

	err = json.Unmarshal(dataJson, request)
	if err != nil {
		logger.Logger.Error("sendRawTransaction", zap.Error(err))
		return nil, 500, err
	}

	err = validator.New().Struct(request)
	if err != nil {
		logger.Logger.Error("sendRawTransaction", zap.Error(err))
		return nil, 500, err
	}

	tx, err := is.Cyclone.GenerateTransactionMessage(request.Message)
	if err != nil {
		logger.Logger.Error("sendRawTransaction", zap.Error(err))
		return nil, 500, err
	}

	err = is.Cyclone.SendTx("/api/gw/v1/transaction/send", tx, response)
	if err != nil {
		logger.Logger.Error("sendRawTransaction", zap.Error(err))
		return nil, 500, err
	}

	return response, 200, nil
}
