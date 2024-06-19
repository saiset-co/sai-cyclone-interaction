package crypto

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/saiset-co/sai-cyclone-interaction/logger"
)

type SaiBTC struct {
	cli  *http.Client
	host string
}

func NewSaiBTCCrypto(host string) *SaiBTC {
	return &SaiBTC{
		cli:  &http.Client{},
		host: host,
	}
}

func (s *SaiBTC) SignMessage(ctx context.Context, payload, pk string) (string, error) {
	var result SignMessageRes

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.host, nil)
	if err != nil {
		logger.Logger.Error("SignMessage", zap.Error(err))
		return "", err
	}
	req.URL.RawQuery = "method=signMessage&p=" + pk + "&message=" + payload
	res, err := s.cli.Do(req)
	if err != nil {
		logger.Logger.Error("SignMessage", zap.Error(err))
		return "", err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Logger.Error("SignMessage", zap.Error(err))
		return "", err
	}

	if res.StatusCode != 200 {
		err = errors.New(string(b))
		logger.Logger.Error("SignMessage", zap.Error(err))
		return "", err
	}

	if err = json.Unmarshal(b, &result); err != nil {
		logger.Logger.Error("SignMessage", zap.Error(err))
		return "", err
	}

	return result.Signature, nil
}
