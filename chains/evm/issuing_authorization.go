package evm

import (
	"encoding/json"

	"github.com/anima-protocol/anima-go/models"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type IssuingAuthorizationEIP712 struct {
	Domain  apitypes.TypedDataDomain    `json:"domain"`
	Message models.IssuingAuthorization `json:"message"`
	Types   apitypes.Types              `json:"types"`
}

func GetIssuingAuthorizationEIP712(challenge []byte, signature string) (*models.IssuingAuthorization, error) {
	authorization := IssuingAuthorizationEIP712{}
	if err := json.Unmarshal(challenge, &authorization); err != nil {
		return nil, err
	}

	valid, err := VerifySignature(authorization.Message.Owner.PublicAddress, challenge, signature)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, err
	}

	return &authorization.Message, nil
}
