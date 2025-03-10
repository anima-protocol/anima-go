package starknetTypedData

import (
	"encoding/json"
	"github.com/NethermindEth/starknet.go/typedData"
	"github.com/anima-protocol/anima-go/crypto"
)

type StarknetAuthorizationMessage struct {
	Hash0 string `json:"hash0"`
	Hash1 string `json:"hash1"`
	Hash2 string `json:"hash2"`
}

func CreateStarknetAuthorizationTypedDataDefinition(
	chainId string,
	message []byte,
) (td *typedData.TypedData, err error) {
	var animaTypes []typedData.TypeDefinition

	domDefs := []typedData.TypeParameter{
		{Name: "name", Type: "felt"},
		{Name: "chainId", Type: "felt"},
		{Name: "version", Type: "felt"},
	}

	animaTypes = append(
		animaTypes, typedData.TypeDefinition{
			Name:       "StarkNetDomain",
			Parameters: domDefs,
		},
	)

	msgDefs := []typedData.TypeParameter{
		{Name: "hash0", Type: "felt"},
		{Name: "hash1", Type: "felt"},
		{Name: "hash2", Type: "felt"},
	}
	animaTypes = append(
		animaTypes, typedData.TypeDefinition{
			Name:       "Message",
			Parameters: msgDefs,
		},
	)

	domain := typedData.Domain{
		Name:    "Anima",
		Version: "1.0.0",
		ChainId: chainId,
	}

	return typedData.NewTypedData(animaTypes, "Message", domain, message)
}

func CreateStarknetAuthorizationTypedDataMessage(data []byte) ([]byte, error) {
	hash := crypto.HashSHA256(data)

	m := StarknetAuthorizationMessage{
		Hash0: hash[:31],
		Hash1: hash[31:62],
		Hash2: hash[62:],
	}

	return json.Marshal(m)
}
