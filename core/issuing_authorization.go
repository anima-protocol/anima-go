package core

import (
	"encoding/base64"
	"encoding/json"

	"github.com/anima-protocol/anima-go/chains"
	cosmos "github.com/anima-protocol/anima-go/chains/cosmos/signature"
	"github.com/anima-protocol/anima-go/chains/elrond"
	"github.com/anima-protocol/anima-go/chains/evm"
	starknet "github.com/anima-protocol/anima-go/chains/starknet/signature"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/protocol"
)

func GetChainSignatureFuncIssuing(authorization *models.IssuingAuthorization) func(string, []byte, string) error {
	switch authorization.Owner.Chain {
	case chains.EVM:
		return evm.VerifyPersonalSignature
	case chains.ELROND:
		return elrond.VerifyPersonalSignature
	case chains.MULTIVERSX:
		return elrond.VerifyPersonalSignature
	case chains.STARKNET:
		return starknet.VerifyPersonalSignature
	case chains.COSMOS:
		return cosmos.VerifyPersonalSignature
	}

	return evm.VerifyPersonalSignature
}

func GetIssuingAuthorization(document *protocol.IssDocument) (*models.IssuingAuthorization, error) {
	encodedContent := document.Authorization.Content
	signature := document.Authorization.Signature

	content, err := base64.StdEncoding.DecodeString(encodedContent)
	if err != nil {
		return nil, err
	}

	authorization := models.IssuingAuthorization{}
	if err := json.Unmarshal(content, &authorization); err != nil {
		return nil, err
	}

	err = GetChainSignatureFuncIssuing(&authorization)(authorization.Owner.PublicAddress, content, signature)
	if err != nil {
		return nil, err
	}

	return &authorization, nil
}
