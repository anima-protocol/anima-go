package starknetTypedData

import (
	"github.com/anima-protocol/anima-go/crypto"
	"github.com/dontpanicdao/caigo"
	"github.com/dontpanicdao/caigo/types"
	"math/big"
)

type StarknetAuthorizationMessage struct {
	Hash0 string
	Hash1 string
	Hash2 string
}

func (m StarknetAuthorizationMessage) FmtDefinitionEncoding(field string) (fmtEnc []*big.Int) {
	if field == "hash0" {
		fmtEnc = append(fmtEnc, types.UTF8StrToBig(m.Hash0))
	} else if field == "hash1" {
		fmtEnc = append(fmtEnc, types.UTF8StrToBig(m.Hash1))
	} else if field == "hash2" {
		fmtEnc = append(fmtEnc, types.UTF8StrToBig(m.Hash2))
	}

	return fmtEnc
}

func CreateStarknetAuthorizationTypedDataDefinition(chainId string) (td caigo.TypedData, err error) {
	animaTypes := make(map[string]caigo.TypeDef)

	domDefs := []caigo.Definition{
		{"name", "felt"},
		{"chainId", "felt"},
		{"version", "felt"},
	}
	animaTypes["StarkNetDomain"] = caigo.TypeDef{Definitions: domDefs}

	msgDefs := []caigo.Definition{
		{"hash0", "felt"},
		{"hash1", "felt"},
		{"hash2", "felt"},
	}
	animaTypes["Message"] = caigo.TypeDef{Definitions: msgDefs}

	domain := caigo.Domain{
		Name:    "Anima",
		Version: "1.0.0",
		ChainId: chainId,
	}

	return caigo.NewTypedData(animaTypes, "Message", domain)
}

func CreateStarknetAuthorizationTypedDataMessage(data []byte) StarknetAuthorizationMessage {
	hash := crypto.HashSHA256(data)

	return StarknetAuthorizationMessage{
		Hash0: hash[:31],
		Hash1: hash[31:62],
		Hash2: hash[62:],
	}
}
