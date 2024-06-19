package internal

import (
	"net/http"

	"github.com/spf13/cast"

	"github.com/saiset-co/sai-cyclone-interaction/internal/crypto"
	"github.com/saiset-co/sai-cyclone-interaction/internal/cyclone"
	saiService "github.com/saiset-co/sai-service/service"
)

type InternalService struct {
	Context *saiService.Context
	client  http.Client
	Cyclone cyclone.Cyclone
}

func (is *InternalService) Init() {
	saiBtc := crypto.NewSaiBTCCrypto(cast.ToString(is.Context.GetConfig("crypto_address", "")))

	is.Cyclone = cyclone.NewCyclone(
		is.Context.Context,
		cast.ToString(is.Context.GetConfig("node_address", "")),
		cast.ToString(is.Context.GetConfig("wallet", "")),
		cast.ToString(is.Context.GetConfig("private", "")),
		saiBtc,
	)
}

func (is *InternalService) Process() {

}
