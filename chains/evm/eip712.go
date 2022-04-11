package evm

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func GetEIP712Message(data []byte) ([]byte, error) {
	signerData := apitypes.TypedData{}
	err := json.Unmarshal(data, &signerData)
	if err != nil {
		return []byte{}, err
	}

	domainSeparator, err := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())
	if err != nil {
		return []byte{}, err
	}

	typedDataHash, err := signerData.HashStruct("Main", signerData.Message)
	if err != nil {
		return []byte{}, err
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	challengeHash := crypto.Keccak256Hash(rawData)
	return challengeHash[:], nil
}
