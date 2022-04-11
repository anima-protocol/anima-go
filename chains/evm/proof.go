package evm

import (
	"bytes"
	"encoding/json"

	"github.com/anima-protocol/anima-go/crypto"
	"github.com/anima-protocol/anima-go/models"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func SignProof(protocol *models.Protocol, content []byte, signingFunc func([]byte) (string, error)) (string, error) {
	message := make(map[string]interface{})

	contentBytes := new(bytes.Buffer)
	err := json.Compact(contentBytes, content)
	if err != nil {
		return "", err
	}

	message["content"] = crypto.Hash(contentBytes.Bytes())

	sigRequest := apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			Name:    models.PROTOCOL_NAME,
			Version: models.PROTOCOL_VERSION,
			ChainId: math.NewHexOrDecimal256(1),
		},
		PrimaryType: "Main",
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
				{
					Name: "version",
					Type: "string",
				},
			},
			"Main": []apitypes.Type{
				{
					Name: "content",
					Type: "string",
				},
			},
		},
		Message: message,
	}

	c, err := json.Marshal(sigRequest)
	if err != nil {
		return "", err
	}

	digest, err := GetEIP712Message(c)
	if err != nil {
		return "", err
	}

	signature, err := signingFunc(digest)
	if err != nil {
		return "", err
	}

	return signature, nil
}
