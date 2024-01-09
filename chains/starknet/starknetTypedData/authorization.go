package starknetTypedData

import (
	"fmt"
	"github.com/NethermindEth/starknet.go/typed"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/anima-protocol/anima-go/crypto"
	"math/big"
	"regexp"
)

type StarknetAuthorizationMessage struct {
	Hash0 string
	Hash1 string
	Hash2 string
}

type BuggedStarknetAuthorizationMessage struct {
	Hash0 string
	Hash1 string
	Hash2 string
}

var digitCheck = regexp.MustCompile(`^[0-9]+$`)

func (m BuggedStarknetAuthorizationMessage) FmtDefinitionEncoding(field string) (fmtEnc []*big.Int) {
	if field == "hash0" {
		if digitCheck.MatchString(m.Hash0) {
			fmtEnc = append(fmtEnc, utils.StrToBig(m.Hash0))
		} else {
			fmtEnc = append(fmtEnc, utils.UTF8StrToBig(m.Hash0))
		}
	} else if field == "hash1" {
		if digitCheck.MatchString(m.Hash1) {
			fmtEnc = append(fmtEnc, utils.StrToBig(m.Hash1))
		} else {
			fmtEnc = append(fmtEnc, utils.UTF8StrToBig(m.Hash1))
		}
	} else if field == "hash2" {
		fmt.Printf("hash2: %s\n", m.Hash2)
		if digitCheck.MatchString(m.Hash2) {
			fmt.Printf("hash2 is digit\n")
			fmtEnc = append(fmtEnc, utils.StrToBig(m.Hash2))
		} else {
			fmtEnc = append(fmtEnc, utils.UTF8StrToBig(m.Hash2))
		}
	}

	return fmtEnc
}

func (m StarknetAuthorizationMessage) FmtDefinitionEncoding(field string) (fmtEnc []*big.Int) {
	if field == "hash0" {
		fmtEnc = append(fmtEnc, utils.UTF8StrToBig(m.Hash0))
	} else if field == "hash1" {
		fmtEnc = append(fmtEnc, utils.UTF8StrToBig(m.Hash1))
	} else if field == "hash2" {
		fmtEnc = append(fmtEnc, utils.UTF8StrToBig(m.Hash2))
	}

	return fmtEnc
}

func CreateStarknetAuthorizationTypedDataDefinition(chainId string) (td typed.TypedData, err error) {
	animaTypes := make(map[string]typed.TypeDef)

	domDefs := []typed.Definition{
		{"name", "felt"},
		{"chainId", "felt"},
		{"version", "felt"},
	}
	animaTypes["StarkNetDomain"] = typed.TypeDef{Definitions: domDefs}

	msgDefs := []typed.Definition{
		{"hash0", "felt"},
		{"hash1", "felt"},
		{"hash2", "felt"},
	}
	animaTypes["Message"] = typed.TypeDef{Definitions: msgDefs}

	domain := typed.Domain{
		Name:    "Anima",
		Version: "1.0.0",
		ChainId: chainId,
	}

	return typed.NewTypedData(animaTypes, "Message", domain)
}

func CreateStarknetAuthorizationTypedDataMessage(data []byte) StarknetAuthorizationMessage {
	hash := crypto.HashSHA256(data)

	return StarknetAuthorizationMessage{
		Hash0: hash[:31],
		Hash1: hash[31:62],
		Hash2: hash[62:],
	}
}

func CreateBuggedStarknetAuthorizationTypedDataMessage(data []byte) BuggedStarknetAuthorizationMessage {
	hash := crypto.HashSHA256(data)

	return BuggedStarknetAuthorizationMessage{
		Hash0: hash[:31],
		Hash1: hash[31:62],
		Hash2: hash[62:],
	}
}
