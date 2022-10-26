package core

import (
	"encoding/base64"
	"encoding/json"

	"github.com/anima-protocol/anima-go/chains"
	elrond "github.com/anima-protocol/anima-go/chains/elrond"
	"github.com/anima-protocol/anima-go/chains/evm"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/protocol"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func GetChainSignatureFuncIssuing(authorization *IssuingAuthorization) func([]byte, string) (*models.IssuingAuthorization, error) {
	switch authorization.Message.Owner.Chain {
	case chains.EVM:
		return evm.GetIssuingAuthorizationEIP712
	case chains.ELROND:
		return elrond.GetIssuingAuthorizationERD712
	}

	return evm.GetIssuingAuthorizationEIP712
}

type IssuingAuthorization struct {
	Domain  apitypes.TypedDataDomain    `json:"domain"`
	Message models.IssuingAuthorization `json:"message"`
	Types   apitypes.Types              `json:"types"`
}

func GetIssuingAuthorization(document *protocol.IssDocument) (*models.IssuingAuthorization, error) {
	encodedContent := document.Authorization.Content
	signature := document.Authorization.Signature

	content, err := base64.StdEncoding.DecodeString(encodedContent)
	if err != nil {
		return nil, err
	}

	authorization := IssuingAuthorization{}
	if err := json.Unmarshal(content, &authorization); err != nil {
		return nil, err
	}

	issuingAuthorization, rErr := GetChainSignatureFuncIssuing(&authorization)(content, signature)
	if rErr != nil {
		return nil, rErr
	}

	return issuingAuthorization, nil
}
