package evm

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

func RecoverAccount(privateKey string) error {
	b, err := hex.DecodeString(privateKey)
	if err != nil {
		return err
	}

	if _, err := crypto.ToECDSA(b); err != nil {
		return err
	}
	return nil
}
