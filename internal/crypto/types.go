package crypto

type ValidateMessageRes struct {
	Address   string `json:"address"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type Secret struct {
	Private string `json:"Private"`
	Public  string `json:"Public"`
	Address string `json:"Address"`
}

type SignMessageRes struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
}
